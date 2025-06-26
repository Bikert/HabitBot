package scenaries

import (
	"HabitMuse/internal/constants"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ScenarioFactory interface {
	GetScenarios() map[string]Scenario
}

type scenarioFactoryImpl struct {
	sessionService session.Service
	userService    users.Service
	botAPI         *tgbotapi.BotAPI
	habitService   habits.Service
}

type Scenario interface {
	Process(sess *session.Session, msg *tgbotapi.Message) error
}

func NewScenarioFactory(sessionService session.Service, userService users.Service, botAPI *tgbotapi.BotAPI, habitService habits.Service) ScenarioFactory {
	return &scenarioFactoryImpl{
		sessionService: sessionService,
		userService:    userService,
		botAPI:         botAPI,
		habitService:   habitService,
	}
}

func (sf scenarioFactoryImpl) GetScenarios() map[string]Scenario {
	return map[string]Scenario{
		constants.Registration.Title: InitRegistration(
			sf.sessionService,
			sf.userService,
			sf.botAPI,
		),
	}
}
