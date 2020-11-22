package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/domain/service"
	fxHandler "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/rest/handler"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/routes"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		service.DefaultModule(),
		fxHandler.AccountModule(),
		fxHandler.TransactionModule(),
		fx.Provide(
			validator.New,
			routes.Routes,
		),
	)
}
