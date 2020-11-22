package viper

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/viper"
	"go.uber.org/fx"
)

func DBModuleConfig() fx.Option {
	return fx.Options(
		fx.Provide(
			viper.DBConfig,
		),
	)
}
