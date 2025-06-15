package appctx

import (
	"HabitMuse/internal/db"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppContext struct {
	SessionService session.SessionService
	UserService    users.Service
	HabitService   habits.Service
	TgBot          *tgBotAPI.BotAPI
}

type BotContext struct {
	AppContext *AppContext
	BotAPI     *tgBotAPI.BotAPI
	Message    *tgBotAPI.Message
	Session    *session.Session
	User       *users.User
}

func BuildApp() *AppContext {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tgToken := os.Getenv("TG_Token")

	fmt.Println("TgToken:", tgToken)
	botAPI, err := tgBotAPI.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bot is worked Token = ", botAPI.Token)

	fmt.Println("connecting to db ...")
	database, err := db.InitDB("HabitMuseDB.db")
	if err != nil {
		fmt.Println("not connected to db", err)
		log.Panic(err)
	}
	fmt.Println("connected to db")

	return &AppContext{
		SessionService: *session.InitSessionService(database),
		UserService:    *users.InitUserService(database),
		HabitService:   *habits.InitHabitService(database),
		TgBot:          botAPI,
	}
}

func (c *BotContext) Send(text string) error {
	_, err := c.BotAPI.Send(
		tgBotAPI.NewMessage(c.Message.Chat.ID, text),
	)
	return err
}
