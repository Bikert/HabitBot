package steps

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/constants"
	"HabitMuse/internal/users"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func verifyUserProfile(ctx *appctx.BotContext) error {
	var newUser users.User
	tgUser := ctx.Message.From
	newUser.UserID = tgUser.ID

	if tgUser.FirstName == "" {
		ctx.Session.PreviousStep = constants.Registration.VerifyProfile
		ctx.Session.NextStep = constants.Registration.FirstNameReceive

		text := fmt.Sprintf(GetResponse("welcome_message_ask_first_name"))
		msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)

		ctx.BotAPI.Send(msg)
		return nil
	}
	newUser.FirstName = tgUser.FirstName

	if tgUser.LastName == "" {
		ctx.Session.PreviousStep = constants.Registration.VerifyProfile
		ctx.Session.NextStep = constants.Registration.LastNameReceive
		text := fmt.Sprintf(GetResponse("welcome_message_ask_last_name"), newUser.FirstName)
		msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
		ctx.BotAPI.Send(msg)
		return nil
	}
	newUser.LastName = tgUser.LastName
	saveUser(newUser, ctx)

	text := fmt.Sprintf(GetResponse("welcome_message_success_registration"), newUser.FirstName)
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)

	ctx.Session.PreviousStep = constants.Registration.VerifyProfile
	ctx.Session.NextStep = constants.MainMenu
	ctx.BotAPI.Send(msg)
	return nil
}

func firsNameReceive(ctx *appctx.BotContext) error {

	firstName := ctx.Message.Text
	tgUser := ctx.Message.From
	if tgUser.LastName == "" {
		ctx.Session.PreviousStep = constants.Registration.FirstNameReceive
		ctx.Session.NextStep = constants.Registration.LastNameReceive
		ctx.Session.Data = map[string]string{
			"first_name": firstName,
		}
		text := fmt.Sprintf(GetResponse("registration_ask_last_name"))
		msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
		ctx.BotAPI.Send(msg)
		return nil
	}

	var newUser users.User
	newUser.UserID = tgUser.ID
	newUser.FirstName = firstName
	newUser.LastName = tgUser.LastName

	saveUser(newUser, ctx)

	text := fmt.Sprintf(GetResponse("success_registration"), newUser.FirstName)
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)

	ctx.Session.PreviousStep = constants.Registration.FirstNameReceive
	ctx.Session.NextStep = constants.MainMenu
	ctx.BotAPI.Send(msg)
	return nil
}

func lastNameReceive(ctx *appctx.BotContext) error {
	tgUser := ctx.Message.From

	var newUser users.User
	newUser.UserID = tgUser.ID
	newUser.FirstName = tgUser.FirstName
	newUser.LastName = ctx.Message.Text
	if ctx.Session.PreviousStep == constants.Registration.FirstNameReceive {
		newUser.FirstName = ctx.Session.Data["first_name"]
	}
	saveUser(newUser, ctx)

	text := fmt.Sprintf(GetResponse("success_registration"), newUser.FirstName)
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)

	ctx.Session.PreviousStep = constants.Registration.LastNameReceive
	ctx.Session.NextStep = constants.MainMenu

	ctx.BotAPI.Send(msg)
	return nil
}

func saveUser(user users.User, ctx *appctx.BotContext) {
	_, err := ctx.AppContext.UserService.SaveOrCreate(user)
	if err != nil {
		msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, "Произошла ошибка попробуйте позже")
		ctx.BotAPI.Send(msg)
		log.Fatal(err)
	}
}
