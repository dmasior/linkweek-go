package repository

import (
	"database/sql"
	"linkweek-go/internal/db/legacy"
	"linkweek-go/internal/models"
	"log"
)

var db *sql.DB

type TopStoryPublication struct {
	PublicationId int
	TopStoryIds   []int
}

func init() {
	db = legacy.Build()
}

func GetItems(offset uint, limit uint) []models.Story {
	rows, err := db.Query("SELECT id, url, title, score, type, FROM_UNIXTIME(time) as time from item order by id asc limit ? offset ?", limit, offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var stories []models.Story

	for rows.Next() {
		var story models.Story
		err = rows.Scan(&story.ID, &story.Url, &story.Title, &story.Score, &story.Type, &story.CreatedAt)

		if err != nil {
			log.Fatal(err)
		}

		stories = append(stories, story)
	}

	return stories
}

func CountItems() int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM item").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

func GetPublications() []models.Publication {
	rows, err := db.Query("SELECT id, date_from, date_to, wykop_entry_id, blog_link, created_at, updated_at FROM publish")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication
		err = rows.Scan(&publication.ID, &publication.DateFrom, &publication.DateTo, &publication.WykopEntryId, &publication.BlogUrl, &publication.CreatedAt, &publication.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		publications = append(publications, publication)
	}

	return publications
}

func GetTopStoryPublications() []TopStoryPublication {
	rows, err := db.Query("SELECT item_id, publish_id FROM top_story ts WHERE ts.publish_id IS NOT NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tsPubMap := map[int][]int{}
	for rows.Next() {
		var pubId, topStoryId int
		err = rows.Scan(&topStoryId, &pubId)
		if err != nil {
			log.Fatal(err)
		}

		tsPubMap[pubId] = append(tsPubMap[pubId], topStoryId)
	}

	topStoryPublications := make([]TopStoryPublication, 0, len(tsPubMap))

	for pubId, tsIds := range tsPubMap {
		topStoryPublications = append(topStoryPublications, TopStoryPublication{PublicationId: pubId, TopStoryIds: tsIds})
	}

	return topStoryPublications
}
