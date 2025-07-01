package scenaries

import (
	"HabitMuse/internal/session"
	"encoding/json"
	"errors"
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

func loadQuestions(filename string) ([]Question, error) {
	data, err := os.ReadFile("resources/questions/" + filename + ".json")
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

func getQuestion(session *session.Session, questionsMap map[string]Question) *Question {
	log.Println("getQuestion started", session.NextStep)
	if session.PreviousStep == "" {
		return getFirstQuestion(questionsMap)
	}
	if session.NextStep == "" {
		return nil
	}
	return getQuestionByID(session.NextStep, questionsMap)
}

func getFirstQuestion(questionsMap map[string]Question) *Question {
	for _, q := range questionsMap {
		if q.Prev == "" {
			return &q
		}
	}
	return nil
}

func getQuestionByID(id string, questionsMap map[string]Question) *Question {
	q, ok := questionsMap[id]
	if !ok {
		return nil
	}
	return &q
}

func ValidateAnswer(q Question, answer string) (bool, error) {
	switch q.Type {
	case "number":
		_, err := strconv.Atoi(answer)
		if err != nil {
			return false, errors.New("❌ Пожалуйста, введите число.")
		}
		return true, nil
	case "choice":
		for _, opt := range q.Options {
			if strings.EqualFold(opt, answer) {
				return true, nil
			}
		}
		return false, errors.New("❌ Неверный ответ. Пожалуйста, выберите один из следующих вариантов: " + strings.Join(q.Options, ", "))
	case "text":
		if len(answer) < 3 {
			return false, errors.New("❌ Пожалуйста, введите текст длиной более 3 символов.")
		}
		return true, nil
	default:
		return false, errors.New("❌ Некорректный ответ. Пожалуйста, попробуйте ещё раз.")
	}
}
