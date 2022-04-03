package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DatabaseDsn string
}

func Provider() Config {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")

	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	return Config{DatabaseDsn: viper.GetString("database.dsn")}
}

var ProviderSet = wire.NewSet(Provider)
