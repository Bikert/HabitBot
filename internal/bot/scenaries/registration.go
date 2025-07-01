package scenaries

import (
	"HabitMuse/internal/bot/utils"
	"HabitMuse/internal/constants"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Registration struct {
	sessionService session.Service
	userService    users.Service
	bot            *tgbotapi.BotAPI
	QuestionsMap   map[string]Question
}

func InitRegistration(sessionService session.Service, userService users.Service, botAPI *tgbotapi.BotAPI) Registration {
	questions, err := loadQuestions("registration")
	questionsMap := make(map[string]Question)
	for _, q := range questions {
		questionsMap[q.ID] = q
	}

	if err != nil {
		return Registration{}
	}
	return Registration{
		sessionService: sessionService,
		userService:    userService,
		bot:            botAPI,
		QuestionsMap:   questionsMap,
	}

}

func (r Registration) StepResolver(sess *session.Session, update *tgbotapi.Update) error {
	message := utils.GetMessage(update)
	q := getQuestion(sess, r.QuestionsMap)
	if q.ID == "last" {
		err, done := r.SaveAnswerToSession(sess, message)
		if done {
			return err
		}
		err = r.saveToDB(sess.Data)
		if err != nil {
			return err
		}
		sess.Scenario = constants.MainMenu
		sess.NextStep = ""
		sess.PreviousStep = ""
		sess.Data = map[string]string{}
	} else {
		if sess.PreviousStep != "" {
			err, done := r.SaveAnswerToSession(sess, message)
			if done {
				return err
			}
		}
		sess.PreviousStep = q.ID
		sess.NextStep = q.Next
	}

	var msg tgbotapi.Chattable
	msg = tgbotapi.NewMessage(sess.UserID, q.Text)
	_, err := r.bot.Send(msg)
	return err

}

func (r Registration) SaveAnswerToSession(sess *session.Session, message *tgbotapi.Message) (error, bool) {
	pq := getQuestionByID(sess.PreviousStep, r.QuestionsMap)
	_, err := ValidateAnswer(*pq, message.Text)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(sess.UserID, err.Error())
		_, err = r.bot.Send(msg)
		return err, true
	}
	sess.Data[sess.PreviousStep] = message.Text
	return nil, false
}

func (r Registration) saveToDB(data map[string]string) error {
	// userService save registration form
	log.Println("Saving data to database")
	for key, value := range data {
		log.Println(key + ": " + value)
	}
	return nil
}
