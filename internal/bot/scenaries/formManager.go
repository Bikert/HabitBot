package scenaries

import (
	"HabitMuse/internal/session"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"` // "text", "number", "choice"
	Next    string   `json:"next"`
	Prev    string   `json:"prev"`
	Options []string `json:"options,omitempty"`
}

type FormQuestion struct {
	questionsMap map[string]Question
}

func loadQuestions(filename string) ([]Question, error) {
	data, err := os.ReadFile("internal/bot/questions/" + filename + ".json")
	if err != nil {
		return nil, err
	}
	var qs []Question
	err = json.Unmarshal(data, &qs)
	if err != nil {
		return nil, err
	}
	return qs, nil
}

func InitFormQuestion(questions []Question) FormQuestion {
	questionsMap := make(map[string]Question)
	for _, q := range questions {
		questionsMap[q.ID] = q
	}
	return FormQuestion{questionsMap: questionsMap}
}

func (f *FormQuestion) isLastQuestion(session session.Session) bool {
	return session.PreviousStep != "" && session.NextStep == ""
}
func (f *FormQuestion) getQuestion(session *session.Session) *Question {
	log.Println("getQuestion started", session.NextStep)
	if session.PreviousStep == "" {
		return f.getFirstQuestion()
	}
	if session.NextStep == "" {
		return nil
	}
	return f.getQuestionByID(session.NextStep)
}

func (f *FormQuestion) getFirstQuestion() *Question {
	for _, q := range f.questionsMap {
		if q.Prev == "" {
			return &q
		}
	}
	return nil
}

func (f *FormQuestion) getQuestionByID(id string) *Question {
	q, ok := f.questionsMap[id]
	if !ok {
		return nil
	}
	return &q
}

func (f *FormQuestion) handleAnswer(session session.Session, answer string) error {
	q := f.getQuestionByID(session.NextStep)
	if q == nil {
		return fmt.Errorf("вопрос не найден")
	}

	//if !validateAnswer(*q, answer) {
	//	return fmt.Errorf("неверный ответ")
	//}

	session.Data[q.ID] = answer

	// Обновляем шаги
	session.PreviousStep = q.ID
	session.NextStep = q.Next

	return nil
}

func (f *FormQuestion) handleBack(session session.Session) error {
	if session.PreviousStep == "" {
		return fmt.Errorf("назад нельзя, это первый вопрос")
	}

	// Берём предыдущий вопрос
	prevQ := f.getQuestionByID(session.PreviousStep)
	if prevQ == nil {
		return fmt.Errorf("предыдущий вопрос не найден")
	}

	// Для перехода назад нам нужен предыдущий шаг от previousStep
	session.NextStep = session.PreviousStep
	session.PreviousStep = prevQ.Prev

	// Можно удалить старый ответ, если хочешь
	delete(session.Data, session.NextStep)

	return nil
}

func ValidateAnswer(q Question, answer string) bool {
	switch q.Type {
	case "number":
		_, err := strconv.Atoi(answer)
		return err == nil
	case "choice":
		for _, opt := range q.Options {
			if strings.EqualFold(opt, answer) {
				return true
			}
		}
		return false
	case "text":
		return len(answer) > 0
	default:
		return false
	}
}
