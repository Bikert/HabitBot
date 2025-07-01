package bot

import (
	"go.uber.org/fx"
)

var Module = fx.Module("bot",
	fx.Provide(NewBot),
	fx.Invoke(RunBot),
)
