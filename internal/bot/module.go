package bot

import (
	"HabitMuse/internal/bot/scenaries"
	"go.uber.org/fx"
)

var Module = fx.Module("bot",
	fx.Provide(
		NewBot,
		NewHandler,
		scenaries.NewScenarioFactory,
	),
	fx.Invoke(RunBot),
)
