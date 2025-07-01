package main

// @title HabitMuse API
// @version 1.0
// @description API для управления привычками и пользователями
// @host localhost:8080

import (
	"HabitMuse/internal/bot"
	"HabitMuse/internal/channels"
	"HabitMuse/internal/db"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/http"
	"HabitMuse/internal/router"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
)

func main() {
	//logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("Error opening log file %v", err)
	//}
	//defer logFile.Close()
	//
	//mw := io.MultiWriter(os.Stdout, logFile)
	//log.SetOutput(mw)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()
	app := fx.New(
		fx.Provide(db.NewDB),
		fx.Provide(channels.NewInitChannels),
		fx.Provide(http.NewHttpServer),
		users.Module,
		habits.Module,
		session.Module,
		bot.Module,

		fx.Provide(
			router.SetupRouter,
			router.NewProtectedGroup,
		),
		fx.Invoke(
			http.RunHttpServer,
		),
	)

	if err := app.Start(ctx); err != nil {
		fmt.Print(err)
	}
	<-app.Done()
}
