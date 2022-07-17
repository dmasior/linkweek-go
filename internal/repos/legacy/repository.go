package repository

import (
	"database/sql"
	"linkweek-go/internal/db/legacy"
	"linkweek-go/internal/models"
	"log"
)

var db *sql.DB

func init() {
	db = legacy.Build()
}

func GetItems(offset uint, limit uint) []models.TopStory {
	rows, err := db.Query("SELECT id, url, title, score, type, FROM_UNIXTIME(time) as time from item order by id asc limit ? offset ?", limit, offset)
	if err != nil {
		log.Fatal(err)
	}

	var stories []models.TopStory

	for rows.Next() {
		var story models.TopStory
		err = rows.Scan(&story.ID, &story.Url, &story.Title, &story.Score, &story.Type, &story.CreatedAt)

		if err != nil {
			log.Fatal(err)
		}

		stories = append(stories, story)
	}

	return stories
}

func GetPublications() []models.Publication {
	// wip get publications
}
