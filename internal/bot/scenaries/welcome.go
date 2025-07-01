package scenaries

import (
	"HabitMuse/internal/constants"
	"HabitMuse/internal/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Welcome struct {
	sessionService session.Service
	api            *tgbotapi.BotAPI
}

func InitWelcome(service session.Service, api *tgbotapi.BotAPI) Welcome {
	return Welcome{
		sessionService: service,
		api:            api,
	}
}
func (s Welcome) StepResolver(sess *session.Session, update *tgbotapi.Update) error {
	log.Println("Welcome step resolver ")
	if update.CallbackQuery != nil {
		callback := update.CallbackQuery
		switch callback.Data {
		case "start_registration":
			sess.Scenario = constants.ScenarioRegistration
			msg := tgbotapi.NewMessage(callback.From.ID, "Отлично, начинаем! 🚀")
			_, err := s.api.Send(msg)
			if err != nil {
				return err
			}

		case "remind_later":
			sess.Scenario = constants.MainMenu
			msg := tgbotapi.NewMessage(callback.From.ID, "Хорошо, напомню позже ⏳")
			_, err := s.api.Send(msg)
			if err != nil {
				return err
			}

		case constants.ScenarioWelcome:
			return s.sendWelcomeMessage(sess)
		}

		return nil
	}
	return s.sendWelcomeMessage(sess)
}

func (s Welcome) sendWelcomeMessage(sess *session.Session) error {
	welcomeMessage := "🎉 **Добро пожаловать в HabitMuse!**\n\nТвой личный спутник в мире дисциплины и полезных привычек.\n\n🧠 Здесь ты начнёшь путь к лучшей версии себя — шаг за шагом, день за днём.\n\n📊 Отслеживай прогресс, строй привычки, оставайся на пути.\n\n💪 Маленькие действия сегодня — великие перемены завтра.\n\n✨ Прежде чем начать — давай немного познакомимся!\nHabitMuse будет помогать тебе становиться лучше каждый день, но для этого ему нужно узнать тебя чуточку ближе.\n\n📝 Ответь на пару коротких вопросов — это поможет подобрать подходящие привычки, напоминания и вдохновляющие советы именно для тебя.\n\nГотова?"

	msg := tgbotapi.NewMessage(sess.UserID, welcomeMessage)

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🚀 Поехали", "start_registration"),
			tgbotapi.NewInlineKeyboardButtonData("⏳ Напомни потом", "remind_later"),
		),
	)
	msg.ReplyMarkup = buttons
	_, err := s.api.Send(msg)
	if err != nil {
		log.Printf("Welcome: Failed to send the message: %v", err)
		return err
	}
	return nil
}
