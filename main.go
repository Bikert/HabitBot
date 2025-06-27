package main

// @title HabitMuse API
// @version 1.0
// @description API для управления привычками и пользователями
// @host localhost:8080

import (
	"HabitMuse/internal/bot"
	"HabitMuse/internal/db"
	"HabitMuse/internal/habits"
	"HabitMuse/internal/http"
	"HabitMuse/internal/router"
	"HabitMuse/internal/session"
	"HabitMuse/internal/users"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file %v", err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()
	app := fx.New(
		fx.Provide(db.NewDB),
		fx.Provide(NewUserRegisteredCh),
		users.Module,
		habits.Module,
		session.Module,
		fx.Provide(
			bot.NewBot,
			bot.NewHandler,
		),
		fx.Provide(router.SetupRouter),

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

func NewUserRegisteredCh() chan int64 {
	return make(chan int64)
}
