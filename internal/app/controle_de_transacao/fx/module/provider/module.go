package provider

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider"
	"go.uber.org/fx"
)

func AccountModule() fx.Option {
	return fx.Options(
		fx.Provide(
			provider.NewAccount,
		),
	)
}

func TransactionModule() fx.Option {
	return fx.Options(
		fx.Provide(
			provider.NewTransaction,
		),
	)
}
