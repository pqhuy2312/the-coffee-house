package models

import (
	"time"
)


type Topic struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Title string    `json:"title"`
	Slug string `json:"slug" gorm:"unique"`
	Tags []Tag `json:"tags" gorm:"foreignKey:TopicId;constraint:OnDelete:CASCADE;"` 
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
