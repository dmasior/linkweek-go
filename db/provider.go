package db

import (
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"linkweek-go/config"
	"log"
)

func Provider(conf config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.DatabaseDsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

var ProviderSet = wire.NewSet(Provider)
