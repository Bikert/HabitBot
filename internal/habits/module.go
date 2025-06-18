package habits

import (
	"go.uber.org/fx"
)

var Module = fx.Module("habits",
	fx.Provide(
		NewRepository,
		NewService,
	),
)
