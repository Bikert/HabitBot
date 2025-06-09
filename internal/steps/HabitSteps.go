package steps

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/habits"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Сценарий для сохранения привычек

func askHabitName(ctx *appctx.BotContext) error {
	fmt.Println("askHabitName started", ctx)
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, GetResponse("askHabitNameMessage"))
	ctx.BotAPI.Send(msg)

	ctx.Session.Scenario = "creatingNewHabit"
	ctx.Session.Step = "receiveAndSaveNewHabit"
	fmt.Println("askHabitName finished")
	return nil
}
func receiveAndSaveNewHabit(ctx *appctx.BotContext) error {
	fmt.Println("receiveAndSaveNewHabit started", ctx.Message.Text)
	habit := &habits.Habit{
		Title:     ctx.Message.Text,
		IsDefault: false,
	}
	ctx.AppContext.HabitService.Save(habit)

	ctx.AppContext.HabitService.SaveHabitByUser(ctx.User.UserID, habit.ID)

	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, fmt.Sprintf(GetResponse("newHabitSavedMessage"), ctx.Message.Text))
	ctx.BotAPI.Send(msg)

	ctx.Session.Scenario = ""
	ctx.Session.Step = ""

	err := MainMenu(ctx)
	fmt.Println("receiveAndSaveNewHabit finished")
	return err
}

// TODO добавить валидацию
