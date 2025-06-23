package bot

import (
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
	// u.AllowedUpdates = []string{"message", "edited_channel_post", "callback_query"}
	updates := bot.GetUpdatesChan(u)
	return updates
}

func RunBot(lc fx.Lifecycle, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("bot started")
			for update := range updates {
				message := getMessage(update)
				MainMenu(message, bot)
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

func MainMenu(message *tgbotapi.Message, api *tgbotapi.BotAPI) error {
	markup := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonWebApp(
				"Open WebApp with keyboard",
				tgbotapi.WebAppInfo{
					URL: os.Getenv("WEB_APP_URL"),
				},
			)))
	keyboardMessage := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: 0,
			ReplyMarkup:      markup,
		},
		DisableWebPagePreview: false,
		Text: "send keyboard",
	}
	keyboardMessage.ReplyMarkup = markup
	_, err := api.Send(keyboardMessage)
	if err != nil {
		return err
	}

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
	inlineMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Открой приложение, не волнуйся сейчас все работает в тестовом режиме на локальной машине соглашайся на все.")
	msg.ReplyMarkup = inlineMarkup
	_, err = api.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func getMessage(update tgbotapi.Update) *tgbotapi.Message {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message
	}
	return update.Message
}
