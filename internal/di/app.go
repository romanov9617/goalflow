package di

import "go.uber.org/fx"

func CreateApp() *fx.App {
	return fx.New(
		fx.WithLogger(NewFxLogger),
		fx.Provide(NewLogger),
		fx.Provide(NewConfig),
		fx.Invoke(NewHTTPServer),
	)
}
