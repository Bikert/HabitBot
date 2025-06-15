package steps

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/constants"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StepFunc func(ctx *appctx.BotContext) error

var stepsMap = map[string]StepFunc{
	constants.MainMenu:                         MainMenu,
	constants.Registration.VerifyProfile:       verifyUserProfile,
	constants.Registration.LastNameReceive:     lastNameReceive,
	constants.Registration.FirstNameReceive:    firsNameReceive,
	constants.HabitCreation.AskTitle:           askHabitName,
	constants.HabitCreation.ReceiveNameAndSave: receiveAndSaveNewHabit,
}

var callBackMap = map[string]StepFunc{
	constants.HabitCreation.AskTitle: askHabitName,
}

func GetStepFuncByCallBack(callback string) StepFunc {
	if fn, ok := callBackMap[callback]; ok {
		return fn
	}
	return Fallback
}

func GetStepFunc(step string) StepFunc {
	if step == "" {
		return MainMenu
	}
	if fn, ok := stepsMap[step]; ok {
		return fn
	}
	return Fallback
}

func MainMenu(ctx *appctx.BotContext) error {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Создать привычку", constants.HabitCreation.AskTitle),
			tgbotapi.NewInlineKeyboardButtonData("Заполнить трекер", "fill_tracker"),
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	text := GetResponse("mainMenuMessage")
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	msg.ReplyMarkup = markup
	ctx.BotAPI.Send(msg)
	ctx.Session.NextStep = ""
	ctx.Session.PreviousStep = constants.MainMenu

	return nil
}

func Fallback(ctx *appctx.BotContext) error {
	text := "Что-то пошло не так, начнём сначала."
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	ctx.BotAPI.Send(msg)
	ctx.Session.NextStep = constants.MainMenu
	return nil
}
