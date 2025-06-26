package scenaries

import (
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Registration struct {
	sessionService      session.Service
	userService         users.Service
	bot                 *tgbotapi.BotAPI
	FormQuestionService FormQuestion
	Questions           []Question
}

func InitRegistration(sessionService session.Service, userService users.Service, botAPI *tgbotapi.BotAPI) Registration {
	questions, err := loadQuestions("registration")
	if err != nil {
		return Registration{}
	}
	fq := InitFormQuestion(questions)
	return Registration{
		sessionService:      sessionService,
		userService:         userService,
		bot:                 botAPI,
		FormQuestionService: fq,
		Questions:           questions,
	}

}

func (r Registration) Process(sess *session.Session, message *tgbotapi.Message) error {

	fq := InitFormQuestion(r.Questions)

	q := fq.getQuestion(sess)
	if q == nil {
		// lastQuestion
		err := r.saveToDB(sess.Data)
		if err != nil {
			return err
		}
		sess.Scenario = ""
		sess.NextStep = ""
		sess.PreviousStep = ""
		sess.Data = map[string]string{}
		return nil
	}

	if sess.PreviousStep != "" {
		sess.Data[sess.PreviousStep] = message.Text
	}

	var msg tgbotapi.Chattable
	msg = tgbotapi.NewMessage(sess.UserID, q.Text)

	_, err := r.bot.Send(msg)
	sess.NextStep = q.Next
	sess.PreviousStep = q.ID
	return err

}

func (r Registration) saveToDB(data map[string]string) error {
	// userService save registration form
	log.Println("Saving data to database")
	for key, value := range data {
		log.Println(key + ": " + value)
	}
	return nil
}
