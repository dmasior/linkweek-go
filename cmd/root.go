package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "linkweek",
		Short: "LinkWeek App",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")
	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	dsn := viper.GetString("database.dsn")

	// wip: handle db
	_, gormErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if gormErr != nil {
		log.Fatal(gormErr)
	}
}
