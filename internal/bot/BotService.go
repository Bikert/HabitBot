package bot

import (
	"HabitMuse/internal/appctx"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"go.uber.org/fx"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot() (*tgbotapi.BotAPI, error) {
	log.Println("creating new bot ... ")
	tgToken := os.Getenv("TG_TOKEN")
	log.Println("tgToken:", tgToken)
	botAPI, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}
	botAPI.Debug = true
	log.Println("bot created")
	return botAPI, nil
}

func NewHandler(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return updates
}

func RunBot(lc fx.Lifecycle, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, userService users.Service, sessionService session.Service, habitService habits.Service) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("bot started")
			for update := range updates {
				userId := getUserId(update)
				message := getMessage(update)
				sess := sessionService.GetOrCreate(userId)

				botCtx := appctx.BotContext{
					SessionService: sessionService,
					UserService:    userService,
					HabitService:   habitService,
					BotAPI:         bot,
					Message:        message,
					Session:        sess,
					UserId:         userId,
				}

				stepFunc := GetStepFunc(sess.NextStep)
				if err := stepFunc(&botCtx); err != nil {
					log.Println("Ошибка в шаге:", err)
				}
				sessionService.Save(botCtx.Session)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping bot...")
			bot.StopReceivingUpdates()
			return nil
		},
	})
	return nil
}

func MainMenu(botCtx *appctx.BotContext) error {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.InlineKeyboardButton{
				Text: "Открыть WebApp",
				WebApp: &tgbotapi.WebAppInfo{
					URL: os.Getenv("WEB_APP_URL"),
				},
			},
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(botCtx.Message.Chat.ID, "Открой приложение, не волнуйся сейчас все работает в тестовом режиме на локальной машине соглашайся на все.")
	msg.ReplyMarkup = markup
	botCtx.BotAPI.Send(msg)
	return nil
}

func getMessage(update tgbotapi.Update) *tgbotapi.Message {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message
	}
	return update.Message
}

func getUserId(update tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	return update.Message.From.ID
}
