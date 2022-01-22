package entities

import "gorm.io/gorm"

type ChapterEntity struct {
	gorm.Model
	ID      string `gorm:"primaryKey;"`
	Name    string `json:"name" gorm:"not null;unique;index"`
	Pages   string `json:"pages" gorm:"not null"`
	MangaID string `json:"manga_id" gorm:"not null"`
}
