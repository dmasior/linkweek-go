package publication

import (
	"linkweek-go/internal/db/orm"
	"linkweek-go/internal/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	Db = orm.Build()
}

func Save(publication ...models.Publication) {
	Db.Save(publication)
}
