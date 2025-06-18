package main

import (
	"HabitMuse/internal/bot"
	"HabitMuse/internal/db"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/http"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()
	app := fx.New(
		fx.Provide(db.NewDB),
		users.Module,
		habits.Module,
		session.Module,
		fx.Provide(
			bot.NewBot,
			bot.NewHandler,
		),
		fx.Invoke(
			http.NewHttpServer,
			bot.RunBot,
		),
	)

	if err := app.Start(ctx); err != nil {
		fmt.Print(err)
	}
	<-app.Done()
}
