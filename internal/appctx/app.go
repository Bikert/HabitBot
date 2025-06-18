package appctx

import (
	"HabitMuse/internal/habits"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotContext struct {
	SessionService session.Service
	UserService    users.Service
	HabitService   habits.Service
	BotAPI         *tgBotAPI.BotAPI
	Message        *tgBotAPI.Message
	Session        *session.Session
	User           *users.User
}
