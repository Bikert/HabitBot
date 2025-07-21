package bot

import (
	"HabitMuse/internal/bot/constants"
	"HabitMuse/internal/bot/scenaries"
	"HabitMuse/internal/bot/utils"
	"HabitMuse/internal/config"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/fx"
	"log"
)

type Bot struct {
	userService     users.Service
	sessionService  session.Service
	habitService    habits.Service
	scenarioFactory scenaries.ScenarioFactory
	api             *tgbotapi.BotAPI
	updateConfig    tgbotapi.UpdateConfig
}

func NewBot(userService users.Service, sessionService session.Service, habitService habits.Service) *Bot {
	log.Println("Initializing bot ... ")
	tgToken := config.Get().TGToken
	botAPI, err := tgbotapi.NewBotAPI(tgToken)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	botAPI.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	scenarioFactory := scenaries.NewScenarioFactory(sessionService, userService, botAPI, habitService)

	log.Println("Initialized bot successfully")
	return &Bot{
		sessionService:  sessionService,
		habitService:    habitService,
		userService:     userService,
		api:             botAPI,
		scenarioFactory: scenarioFactory,
		updateConfig:    updateConfig,
	}
}

func RunBot(lc fx.Lifecycle, bot *Bot) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Bot starting ...")
			bot.onStart()
			log.Println("Bot started successfully")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Bot stopping ...")
			bot.api.StopReceivingUpdates()
			log.Println("Bot stopped successfully...")
			return nil
		},
	})
	return nil
}

func (bot *Bot) onStart() {
	go func() {
		updates := bot.api.GetUpdatesChan(bot.updateConfig)
		for update := range updates {
			log.Println("++++++++Get NEW update ++++++++++++")
			log.Println("update = ", update)
			tgUser := utils.GetUserId(&update)
			user, err := bot.userService.GetOrCreateUser(*tgUser)
			if err != nil {
				log.Println(err)
				bot.sendErrorMessage(update, err)
				continue
			}
			message := utils.GetMessage(&update)
			sess := bot.sessionService.GetOrCreateSessionForUser(user.UserID)

			if sess.Scenario == constants.ScenarioWelcome {
				bot.sendWelcomeMessage(update, message)
				sess.Scenario = constants.ScenarioDefault
				bot.sessionService.Save(*sess)
				continue
			}

			bot.sendDefaultMessage(message)

			//if message.IsCommand() {
			//	bot.commandStepResolver(update, sess)
			//	continue
			//}
			//
			//if update.CallbackQuery != nil {
			//	bot.callbackStepResolver(sess, update)
			//	continue
			//}
			//bot.scenarioResolver(sess, update)
		}
	}()
}

func (bot *Bot) commandStepResolver(update tgbotapi.Update, sess *session.Session) {
	message := utils.GetMessage(&update)
	log.Println("message is command ", message.Command())
	switch message.Command() {
	case constants.OpenApp:
		err := bot.sendOpenAppButton(update)
		if err != nil {
			log.Println(err)
		}
	case constants.MainMenu:
		err := bot.mainMenu(update, sess)
		if err != nil {
			log.Println(err)
		}
	}
}

// If a callback is received from MainMenu buttons, we switch the scenario and pass it to the scenarioResolver.
func (bot *Bot) callbackStepResolver(sess *session.Session, update tgbotapi.Update) {
	log.Println("Calling callback step resolver callback data = ", update.CallbackQuery.Data)
	callback := update.CallbackQuery
	err := utils.СonfirmPressAndHideButtons(bot.api, callback)
	if err != nil {
		return
	}

	if sess.Scenario == constants.MainMenu {
		log.Println("setScenario to ", callback.Data)
		sess.Scenario = callback.Data
		sess.NextStep = ""
		sess.PreviousStep = ""
	}
	bot.scenarioResolver(sess, update)
}

func (bot *Bot) scenarioResolver(sess *session.Session, update tgbotapi.Update) {
	scenarios := bot.scenarioFactory.GetScenarios()
	currentScenario := ""
	for {
		if sess.Scenario == "" || sess.Scenario == constants.MainMenu {
			err := bot.mainMenu(update, sess)
			if err != nil {
				log.Println(err)
			}
			break
		}
		if currentScenario == sess.Scenario {
			break
		}
		currentScenario = sess.Scenario

		if scenario, ok := scenarios[sess.Scenario]; ok {
			if err := scenario.StepResolver(sess, &update); err != nil {
				log.Println("Ошибка в сценарии:", err)
				bot.sendErrorMessage(update, err)
			}
			bot.sessionService.Save(*sess)
		} else {
			str := fmt.Sprintf("Сценарий не найден: %s", sess.Scenario)
			log.Println(str)
			bot.sendErrorMessage(update, errors.New(str))
			break
		}
	}
}

func (bot *Bot) sendOpenAppButton(update tgbotapi.Update) error {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.InlineKeyboardButton{
				Text: "Открыть WebApp",
				WebApp: &tgbotapi.WebAppInfo{
					URL: config.Get().WebBaseUrl,
				},
			},
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Открой приложение, не волнуйся сейчас все работает в тестовом режиме на локальной машине соглашайся на все.")
	msg.ReplyMarkup = markup
	_, err := bot.api.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) mainMenu(update tgbotapi.Update, sess *session.Session) error {
	sess.Scenario = constants.MainMenu
	sess.NextStep = ""
	sess.PreviousStep = ""
	sess.Data = map[string]string{}

	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("🚀 Заполнить Анкету Регистрации", constants.ScenarioRegistration),
			tgbotapi.NewInlineKeyboardButtonData("🚀 Отправить приветственное сообщение", constants.ScenarioWelcome),
		},
		{
			tgbotapi.InlineKeyboardButton{
				Text: "Затрекать привычку",
				WebApp: &tgbotapi.WebAppInfo{
					URL: config.Get().WebBaseUrl,
				},
			},
			tgbotapi.InlineKeyboardButton{
				Text: "Пора измеряться!",
				WebApp: &tgbotapi.WebAppInfo{
					URL: config.Get().WebBaseUrl + "body-measurements",
				},
			},
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(utils.GetMessage(&update).Chat.ID, "Вот что я умею")
	msg.ReplyMarkup = markup
	_, err := bot.api.Send(msg)
	if err != nil {
		return err
	}
	bot.sessionService.Save(*sess)
	return nil
}

func (bot *Bot) sendErrorMessage(update tgbotapi.Update, err error) {
	str := fmt.Sprintf("😕 Упс! Что-то пошло не так: %s. Попробуй ещё раз чуть позже!", err.Error())
	msg := tgbotapi.NewMessage(utils.GetMessage(&update).Chat.ID, str)
	_, err = bot.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (bot *Bot) sendWelcomeMessage(update tgbotapi.Update, message *tgbotapi.Message) {
	text := "💁‍♀️ Привет, богиня самодисциплины (ну или на пути к ней)!\n\nТы только что сделала первый шаг к телу мечты — или хотя бы к тому, чтобы не есть чизкейк каждый день 🧁😉\n\n👇 Жми кнопку ниже — там тебя ждёт наше уютное приложение, где ты сможешь следить за своими привычками и становиться лучше день за днём.\n\n⚠️ Приложение ещё в разработке, так что если что-то покажется странным — просто напиши всё сюда: @BikertE 💌\n\nГотова? Ведь красота требует не жертв, а дисциплины — мягкой, но регулярной ✨"
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.InlineKeyboardButton{
				Text: "Трекать, а не страдать 🧘‍♀️",
				WebApp: &tgbotapi.WebAppInfo{
					URL: config.Get().WebBaseUrl,
				},
			},
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyMarkup = markup
	bot.api.Send(msg)
}

func (bot *Bot) sendDefaultMessage(message *tgbotapi.Message) {
	text := "🤷‍♀️ Ой! Кажется, я не поняла, что ты хотела сказать.\nМожет, это секретная команда, которую ещё не придумали? 🕵️‍♀️✨\n\nЕсли что-то не работает или есть идея получше — пиши сюда: @BikertE 💌\nА пока можешь просто нажать кнопку ниже и продолжить путь к супер-версии себя 💪💖"
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.InlineKeyboardButton{
				Text: "К лучшей версии себя 💖",
				WebApp: &tgbotapi.WebAppInfo{
					URL: config.Get().WebBaseUrl,
				},
			},
		},
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyMarkup = markup
	bot.api.Send(msg)

}
