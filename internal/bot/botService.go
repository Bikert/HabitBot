package bot

import (
	"HabitMuse/internal/bot/scenaries"
	"HabitMuse/internal/bot/utils"
	"HabitMuse/internal/constants"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"errors"
	"fmt"
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
			go func() {
				log.Println("bot started")
				for update := range updates {
					log.Println("++++++++Get NEW update ++++++++++++")
					log.Println("update = ", update)
					tgUser := utils.GetUserId(&update)
					user, err := userService.GetOrCreateUser(*tgUser)
					if err != nil {
						log.Println(err)
						SendErrorMessage(update, bot, err)
						continue
					}
					message := utils.GetMessage(&update)
					sess := sessionService.GetOrCreateSessionForUser(user.UserID)

					if message.IsCommand() {
						log.Println("++++++++ COMMAND++++++++++++")
						CommandStepResolver(update, bot)
						log.Println("++++++++ COMMAND  finished +++++++++")
						continue
					}

					if update.CallbackQuery != nil {
						log.Println("++++++++ CALLBACKQUERY ++++++++++++")
						CallbackStepResolver(sess, scenarioFactory, update, sessionService, bot)
						continue
					}
					runScenario(sess, scenarioFactory, update, sessionService, bot)
				}
			}()
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

func CommandStepResolver(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	message := utils.GetMessage(&update)
	log.Println("message is command ", message.Command())
	switch message.Command() {
	case constants.OpenApp:
		err := SendOpenAppButton(update, bot)
		if err != nil {
			log.Println(err)
		}
	case constants.MainMenu:
		err := mainMenu(update, bot)
		if err != nil {
			log.Println(err)
		}
	}
}

func CallbackStepResolver(sess *session.Session, scenarioFactory scenaries.ScenarioFactory, update tgbotapi.Update, sessionService session.Service, bot *tgbotapi.BotAPI) {
	log.Println("callback query", update.CallbackQuery)
	callbackData := update.CallbackQuery.Data
	callbackConfig := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	_, err := bot.Request(callbackConfig)
	if err != nil {
		log.Println(err)
	}

	switch callbackData {
	case constants.CallbackRegistration:
		sess.Scenario = constants.ScenarioRegistration
		sess.NextStep = ""
		sess.PreviousStep = ""
	case constants.CallbackSendWelcomeMessage:
		sess.Scenario = constants.ScenarioWelcome
		sess.NextStep = ""
		sess.PreviousStep = ""
	}
	runScenario(sess, scenarioFactory, update, sessionService, bot)
}

func runScenario(sess *session.Session, scenarioFactory scenaries.ScenarioFactory, update tgbotapi.Update, sessionService session.Service, bot *tgbotapi.BotAPI) {
	if sess.Scenario == "" || sess.Scenario == constants.MainMenu {
		err := mainMenu(update, bot)
		if err != nil {
			return
		}
		return
	}
	scenarios := scenarioFactory.GetScenarios()
	currentScenario := ""
	for {
		if sess.Scenario == "" || currentScenario == sess.Scenario {
			break
		}
		currentScenario = sess.Scenario

		if scenario, ok := scenarios[sess.Scenario]; ok {
			if err := scenario.StepResolver(sess, &update); err != nil {
				log.Println("Ошибка в сценарии:", err)
			}
			sessionService.Save(*sess)
		} else {
			str := fmt.Sprintf("Сценарий не найден: %s", sess.Scenario)
			log.Println(str)
			SendErrorMessage(update, bot, errors.New(str))
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

func mainMenu(update tgbotapi.Update, api *tgbotapi.BotAPI) error {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("🚀 Заполнить Анкету Регистрации", constants.CallbackRegistration),
			tgbotapi.NewInlineKeyboardButtonData("🚀 Привет", constants.CallbackSendWelcomeMessage),
		},
		{
			tgbotapi.InlineKeyboardButton{
				Text: "Затрекать привычку",
				WebApp: &tgbotapi.WebAppInfo{
					URL: os.Getenv("WEB_APP_URL"),
				},
			},
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(utils.GetMessage(&update).Chat.ID, "Вот что я умею")
	msg.ReplyMarkup = markup
	_, err := api.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func SendErrorMessage(update tgbotapi.Update, api *tgbotapi.BotAPI, err error) {
	str := fmt.Sprintf("😕 Упс! Что-то пошло не так: %s. Попробуй ещё раз чуть позже!", err.Error())
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
	_, err = api.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
