package bot

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/steps"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(app *appctx.AppContext) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := app.TgBot.GetUpdatesChan(updateConfig)
	for update := range updates {
		ifCallBack := update.CallbackQuery != nil

		userId := getUser(update).ID
		user, _ := app.UserService.Get(userId)

		sess := app.SessionService.GetOrCreate(userId)
		message := getMessage(update)

		botCtx := &appctx.BotContext{
			AppContext: app,
			BotAPI:     app.TgBot,
			Message:    message,
			Session:    sess,
			User:       &user,
		}

		var stepFunc steps.StepFunc
		if ifCallBack {
			stepFunc = steps.GetStepFuncByCallBack(update.CallbackQuery.Data)
		} else {
			stepFunc = steps.GetStepFunc(sess.NextStep)
		}

		if err := stepFunc(botCtx); err != nil {
			log.Println("Ошибка в шаге:", err)
		}
		app.SessionService.Save(botCtx.Session)
	}

}

func getMessage(update tgbotapi.Update) *tgbotapi.Message {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message
	}
	return update.Message
}

func getUser(update tgbotapi.Update) *tgbotapi.User {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From
	}
	return update.Message.From

}
