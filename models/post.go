package models

import (
	"time"
)


type Post struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Slug string `json:"slug" gorm:"unique"`
	Content string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	TagId int `json:"tagId"`
	Tag Tag `json:"tag" gorm:"foreignKey:TagId;constraint:OnDelete:CASCADE;"`
	AuthorId int
	Author User `json:"author" gorm:"foreignKey:AuthorId;constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
