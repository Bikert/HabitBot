package habits

import (
	"HabitMuse/internal/abstractions"
	"go.uber.org/fx"
)

var Module = fx.Module("habits",
	fx.Provide(
		NewRepository,
		NewService,
		fx.Annotate(
			NewHandler,
			fx.As(new(abstractions.ApiHandler)),
			fx.ResultTags(`group:"apiHandlers"`),
		),
	),
	fx.Invoke(
		UserRegistrationListener,
	),
)
