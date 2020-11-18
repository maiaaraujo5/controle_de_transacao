package runner

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/rest"
	"go.uber.org/fx"
)

func RunApplication() error {
	return fx.New(
		rest.Module(),
		fx.Provide(context.Background),
		fx.Invoke(start),
	).Start(context.Background())
}

func start(lifecycle fx.Lifecycle, e *echo.Echo) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Info("starting...")
				return e.Start(":8080")
			},
			OnStop: func(ctx context.Context) error {
				log.Info("stopping...")
				return e.Close()
			},
		},
	)
}
