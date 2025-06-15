package steps

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/constants"
	"HabitMuse/internal/habits"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Сценарий для сохранения привычек

func askHabitName(ctx *appctx.BotContext) error {
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, GetResponse("askHabitNameMessage"))
	ctx.BotAPI.Send(msg)

	ctx.Session.PreviousStep = constants.HabitCreation.AskTitle
	ctx.Session.NextStep = constants.HabitCreation.ReceiveNameAndSave
	return nil
}
func receiveAndSaveNewHabit(ctx *appctx.BotContext) error {
	habit := &habits.Habit{
		Title:     ctx.Message.Text,
		IsDefault: false,
	}
	ctx.AppContext.HabitService.Save(habit)
	ctx.AppContext.HabitService.SaveHabitByUser(ctx.User.UserID, habit.ID)

	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, fmt.Sprintf(GetResponse("newHabitSavedMessage"), ctx.Message.Text))
	ctx.BotAPI.Send(msg)

	ctx.Session.PreviousStep = constants.HabitCreation.ReceiveNameAndSave
	ctx.Session.NextStep = constants.MainMenu

	err := MainMenu(ctx)
	return err
}

// TODO добавить валидацию
