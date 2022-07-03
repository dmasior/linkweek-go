package db

import (
	"linkweek-go/config"
	"linkweek-go/db/model"
	"log"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Provider(conf config.Config) *gorm.DB {
	db, connErr := gorm.Open(postgres.Open(conf.DatabaseDsn), &gorm.Config{})

	if connErr != nil {
		log.Fatal(connErr)
	}

	migrationErr := db.AutoMigrate(model.TopStory{}) // todo: handle migrations
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	return db
}

var ProviderSet = wire.NewSet(Provider)
