package main

// @title HabitMuse API
// @version 1.0
// @description API для управления привычками и пользователями
// @host localhost:8080

import (
	"HabitMuse/internal/body_metrics"
	"HabitMuse/internal/bot"
	"HabitMuse/internal/channels"
	"HabitMuse/internal/config"
	"HabitMuse/internal/db"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/http"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	config.Init()

	ctx := context.Background()
	app := fx.New(
		fx.Provide(db.NewDB),
		fx.Provide(channels.NewInitChannels),
		fx.Provide(
			http.NewHttpServer,
			fx.Annotate(
				http.NewEngine,
				fx.ParamTags(`group:"apiHandlers"`),
			),
		),
		users.Module,
		habits.Module,
		body_metrics.Module,
		session.Module,
		bot.Module,

		fx.Invoke(
			http.RunHttpServer,
		),
	)

	if err := app.Start(ctx); err != nil {
		fmt.Print(err)
	}
	<-app.Done()
}
