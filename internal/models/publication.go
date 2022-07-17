package models

import "time"

type Publication struct {
	ID           uint       `gorm:"primaryKey"`
	DateFrom     time.Time  `gorm:"not null"`
	DateTo       time.Time  `gorm:"not null"`
	TopStories   []TopStory `gorm:"many2many:publication_topstories;"`
	WykopEntryId uint
	BlogUrl      string
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}
