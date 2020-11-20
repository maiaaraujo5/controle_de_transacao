package handler

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/handler"
	"go.uber.org/fx"
)

func AccountModule() fx.Option {
	return fx.Options(
		service.AccountModule(),
		fx.Provide(
			handler.NewCreateAccount,
			handler.NewRecoverAccount,
		),
	)
}

func TransactionModule() fx.Option {
	return fx.Options(
		service.TransactionModule(),
		fx.Provide(
			handler.NewCreateTransaction,
			handler.NewRecoverTransaction,
		),
	)
}
