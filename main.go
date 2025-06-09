package main

import (
	"HabitMuse/internal/appctx"
	Bot "HabitMuse/internal/bot"
)

func main() {
	app := appctx.BuildApp()
	Bot.Start(app)
}
