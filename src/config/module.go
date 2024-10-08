package config

import (
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(LoadConfig),
	fx.Invoke(func(c *ChatBotConfig) {

	}),
)
