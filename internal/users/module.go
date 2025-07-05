package users

import (
	"HabitMuse/internal/abstractions"
	"go.uber.org/fx"
)

var Module = fx.Module("users",
	fx.Provide(
		NewRepository,
		NewService,
		fx.Annotate(
			NewHandler,
			fx.As(new(abstractions.ApiHandler)),
			fx.ResultTags(`group:"apiHandlers"`),
		),
	),
)
