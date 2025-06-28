package bot

import (
	"HabitMuse/internal/bot/scenaries"
	"HabitMuse/internal/bot/utils"
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
	// u.AllowedUpdates = []string{"message", "edited_channel_post", "callback_query"}
	updates := bot.GetUpdatesChan(u)
	return updates
}

func RunBot(lc fx.Lifecycle, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, userService users.Service, sessionService session.Service, habitService habits.Service, scenarioFactory scenaries.ScenarioFactory) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("bot starting")
			go func() {
				for update := range updates {
					tgUser := utils.GetUserId(&update)
					user, err := userService.GetOrCreateUser(*tgUser)
					if err != nil {
						log.Println(err)
					}
					message := utils.GetMessage(&update)
					sess := sessionService.GetOrCreate(user.UserID)

					switch message.Command() {
					case "open":
						err := SendOpenAppButton(update, bot)
						if err != nil {
							log.Println(err)
						}
					default:
						runScenario(sess, scenarioFactory, update, sessionService, bot)
					}
				}
			}()
			log.Println("bot started")
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

func runScenario(sess session.Session, scenarioFactory scenaries.ScenarioFactory, update tgbotapi.Update, sessionService session.Service, bot *tgbotapi.BotAPI) {
	scenarios := scenarioFactory.GetScenarios()
	currentScenario := ""
	for {
		if sess.Scenario == "" || currentScenario == sess.Scenario {
			break
		}
		currentScenario = sess.Scenario

		if scenario, ok := scenarios[sess.Scenario]; ok {
			if err := scenario.StepResolver(&sess, &update); err != nil {
				log.Println("Ошибка в сценарии:", err)
			}
			sessionService.Save(sess)
		} else {
			log.Println("Сценарий не найден:", sess.Scenario)
			err := fallback(update, bot)
			if err != nil {
				return
			}
			break
		}
	}
}

func SendOpenAppButton(update tgbotapi.Update, api *tgbotapi.BotAPI) error {
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Открой приложение, не волнуйся сейчас все работает в тестовом режиме на локальной машине соглашайся на все.")
	msg.ReplyMarkup = markup
	_, err := api.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func fallback(update tgbotapi.Update, api *tgbotapi.BotAPI) error {
	text := "Что-то пошло не так, начнём сначала."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	api.Send(msg)
	return nil
}
