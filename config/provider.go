package config

import (
	"log"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDsn       string
	LegacyDatabaseDsn string
}

func Provider() Config {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")

	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	return Config{
		DatabaseDsn:       viper.GetString("database.dsn"),
		LegacyDatabaseDsn: viper.GetString("database.legacy.dsn"),
	}
}

var ProviderSet = wire.NewSet(Provider)
