package logs

import (
	"go.uber.org/fx"
)

var Module = fx.Module("logs",
	fx.Provide(NewConsoleLogger),
	fx.Invoke(func(c *ConsoleLogger) {

	}),
)
