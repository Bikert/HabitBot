package scenaries

import (
	"HabitMuse/internal/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BodyParamsActualizationScenario struct {
	sessionService session.Service
	api            *tgbotapi.BotAPI
}

func InitBodyParamsActualizationScenario(service session.Service, api *tgbotapi.BotAPI) BodyParamsActualizationScenario {
	return BodyParamsActualizationScenario{
		sessionService: service,
		api:            api,
	}
}
