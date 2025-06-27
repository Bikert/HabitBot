package session

import (
	"go.uber.org/fx"
)

var Module = fx.Module("session",
	fx.Provide(
		NewRepository,
		NewService,
	),
	fx.Invoke(UserRegistrationListener),
)
