package bot

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/bot/scenaries"
	"HabitMuse/internal/constants"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StepFunc func(ctx *appctx.BotContext) error

var stepsMap = map[string]StepFunc{
	constants.Registration.Title: scenaries.Registration,
}

func GetStepFunc(step string) StepFunc {
	if step == "" {
		return Fallback
	}
	if fn, ok := stepsMap[step]; ok {
		return fn
	}
	return Fallback
}

var callBackMap = map[string]StepFunc{}

func GetStepFuncByCallBack(callback string) StepFunc {
	if fn, ok := callBackMap[callback]; ok {
		return fn
	}
	return Fallback
}

func Fallback(ctx *appctx.BotContext) error {
	text := "Что-то пошло не так, начнём сначала."
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	ctx.BotAPI.Send(msg)
	ctx.Session.NextStep = constants.MainMenu
	return nil
}
