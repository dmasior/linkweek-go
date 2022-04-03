package model

import "time"

type TopStory struct {
	ID        uint      `gorm:"primaryKey"`
	Url       string    `gorm:"size:512;not null"`
	Title     string    `gorm:"size:128;not null"`
	Score     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
