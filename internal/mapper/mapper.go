package mapper

import (
	"linkweek-go/internal/fetch"
	"linkweek-go/internal/models"
	"time"
)

func ItemToDBModels(items []fetch.Item) []models.Story {
	r := make([]models.Story, 0, len(items))

	for _, i := range items {
		r = append(r, models.Story{
			ID:        i.Id,
			Url:       i.Url,
			Title:     i.Title,
			Score:     i.Score,
			Type:      i.Type,
			CreatedAt: time.Unix(int64(i.Time), 0),
		})
	}

	return r
}
