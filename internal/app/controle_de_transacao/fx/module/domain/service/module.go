package service

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/provider"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/provider/postgres"
	"go.uber.org/fx"
)

func DefaultModule() fx.Option {
	return fx.Options(
		postgres.Module(),
	)
}

func AccountModule() fx.Option {
	return fx.Options(
		provider.AccountModule(),
		fx.Provide(
			service.NewCreateAccount,
			service.NewRecoverAccount,
		),
	)
}

func TransactionModule() fx.Option {
	return fx.Options(
		provider.TransactionModule(),
		fx.Provide(
			service.NewCreateTransaction,
			service.NewRecoverTransaction,
		),
	)
}
