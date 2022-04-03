package repo

import (
	"gorm.io/gorm"
	"linkweek-go/db"
	"linkweek-go/db/model"
	"linkweek-go/fetch"
	"log"
	"time"
)

var (
	Db *gorm.DB
)

func init() {
	Db = db.Build()
}

func Persist(items []fetch.Item) {
	topStories := make([]model.TopStory, len(items))
	for i, el := range items {
		topStories[i] = createTopStoryFromItem(el)
	}

	err := Db.AutoMigrate(model.TopStory{})
	if err != nil {
		log.Fatal(err)
	}

	Db.Save(topStories)
}

func createTopStoryFromItem(i fetch.Item) model.TopStory {
	return model.TopStory{
		ID:        i.Id,
		Url:       i.Url,
		Title:     i.Title,
		Score:     i.Score,
		CreatedAt: time.Time{},
	}
}
