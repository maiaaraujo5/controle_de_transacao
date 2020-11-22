package viper

import (
	"fmt"
	"github.com/labstack/gommon/log"
	postgresProvider "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider/postgre/config"
	"github.com/spf13/viper"
	"os"
)

func DBConfig() *postgresProvider.DBConfig {
	environment := getEnvironment()
	log.Infof("getting %s configuration", environment)

	viper.SetConfigName(getEnvironment())
	viper.AddConfigPath("configs/")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	return &postgresProvider.DBConfig{
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
		Adrr:     viper.GetString("db.addr"),
	}
}

func getEnvironment() string {
	value := os.Getenv("environment")
	if value != "" {
		return value
	}
	return "development"
}
