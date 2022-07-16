package orm

import (
	"linkweek-go/config"
	"linkweek-go/internal/models"
	"log"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Provider(conf config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.DatabaseDsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(models.TopStory{}, models.Publication{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

var ProviderSet = wire.NewSet(Provider)
