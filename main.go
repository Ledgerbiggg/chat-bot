package main

import (
	"chat-bot/src/config"
	"chat-bot/src/logs"
	"chat-bot/src/services"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		services.Module,
		logs.Module,
	)
	app.Run()
}
