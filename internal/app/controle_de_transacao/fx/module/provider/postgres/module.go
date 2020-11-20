package postgres

import (
	"github.com/go-pg/pg/v10"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewDBConn,
		),
	)
}

func NewDBConn() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "docker",
		Database: "pismo",
		Addr:     "localhost:5432",
		PoolSize: 5,
	}

	return pg.Connect(options)
}
