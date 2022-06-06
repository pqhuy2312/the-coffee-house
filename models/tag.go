package models

import (
	"time"
)


type Tag struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	TopicId int `json:"topicId"`
	Title string `json:"title"`
	Slug string `json:"slug" gorm:"unique"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
