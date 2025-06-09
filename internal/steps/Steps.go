package steps

import (
	"HabitMuse/internal/appctx"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StepFunc func(ctx *appctx.BotContext) error

var defaultStep = "welcome"

var scenarios = map[string]map[string]StepFunc{
	"welcome": {
		"welcome":      Welcome,
		"regestration": newUser,
	},
	"creatingNewHabit": {
		"askHabitName":           askHabitName,
		"receiveAndSaveNewHabit": receiveAndSaveNewHabit,
	},
}

var callBackMap = map[string]StepFunc{
	"create_habit": askHabitName,
}

func GetStepFuncByCallBack(callback string) StepFunc {
	if fn, ok := callBackMap[callback]; ok {
		return fn
	}
	return Fallback
}

func GetStepFunc(scenario, step string) StepFunc {
	stepsMap, ok := scenarios[scenario]
	if !ok {
		return Welcome
	}
	if step == "" {
		return stepsMap[defaultStep]
	}
	if fn, ok := stepsMap[step]; ok {
		return fn
	}
	return Welcome
}

func Welcome(ctx *appctx.BotContext) error {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Создать привычку", "create_habit"),
			tgbotapi.NewInlineKeyboardButtonData("Заполнить трекер", "fill_tracker"),
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	text := fmt.Sprintf(GetResponse("welcomeMessage"), ctx.Message.From.FirstName)
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	msg.ReplyMarkup = markup
	ctx.BotAPI.Send(msg)
	ctx.Session.Step = ""
	return nil
}

func MainMenu(ctx *appctx.BotContext) error {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Создать привычку", "create_habit"),
			tgbotapi.NewInlineKeyboardButtonData("Заполнить трекер", "fill_tracker"),
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	text := GetResponse("mainMenuMessage")
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	msg.ReplyMarkup = markup
	ctx.BotAPI.Send(msg)
	ctx.Session.Step = ""

	return nil
}

func Fallback(ctx *appctx.BotContext) error {
	text := "Что-то пошло не так, начнём сначала."
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	ctx.BotAPI.Send(msg)
	ctx.Session.Scenario = "registration"
	ctx.Session.Step = "greeting"
	return nil
}
