package legacy

import (
	"database/sql"
	"linkweek-go/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

func Provider(conf config.Config) *sql.DB {
	db, err := sql.Open("mysql", conf.LegacyDatabaseDsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

var ProviderSet = wire.NewSet(Provider)
