package scenaries

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/constants"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Registration(ctx *appctx.BotContext) error {
	if ctx.Session.NextStep == constants.Registration.Title {
		ctx.Session.NextStep = "name"
	}

	questions, err := loadQuestions("registration")
	if err != nil {
		return err
	}
	fq := InitFormQuestion(questions)

	for _, question := range questions {
		fmt.Printf("Question: %v\n", question)
	}

	q := fq.getQuestionByID(ctx.Session.NextStep)

	if q.Prev != "" {
		// TODO Add save answer to data in session
	}

	if ctx.Session.NextStep == "" {
		// Анкета закончена
		msg := tgbotapi.NewMessage(ctx.UserId, "Анкета завершена! Спасибо за ответы.")
		ctx.BotAPI.Send(msg)
		return nil
	}

	var msg tgbotapi.Chattable
	msg = tgbotapi.NewMessage(ctx.UserId, q.Text)

	_, err = ctx.BotAPI.Send(msg)

	ctx.Session.NextStep = q.Next
	return err

}
