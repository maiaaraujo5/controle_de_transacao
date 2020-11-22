package postgres

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/viper"
	postgreconfig "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider/postgre/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		viper.DBModuleConfig(),
		fx.Provide(
			postgreconfig.NewDBConn,
		),
	)
}
