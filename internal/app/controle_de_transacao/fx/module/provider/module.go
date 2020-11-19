package provider

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider/postgre/dao"
	"go.uber.org/fx"
)

func AccountModule() fx.Option {
	return fx.Options(
		fx.Provide(
			dao.NewAccount,
		),
	)
}

func TransactionModule() fx.Option {
	return fx.Options(
		fx.Provide(
			dao.NewTransaction,
		),
	)
}
