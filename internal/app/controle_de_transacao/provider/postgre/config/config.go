package config

import "github.com/go-pg/pg/v10"

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
