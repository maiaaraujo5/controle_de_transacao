package config

import "github.com/go-pg/pg/v10"

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Adrr     string `yaml:"adrr"`
}

func NewDBConn(config *DBConfig) *pg.DB {
	options := &pg.Options{
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
		Addr:     config.Adrr,
		PoolSize: 5,
	}

	return pg.Connect(options)
}
