package models

import (
	"time"
)


type Category struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Slug string `json:"slug" gorm:"unique"`
	ParentId 	*int    `json:"parentId"`
	Children []Category `json:"children" gorm:"foreignkey:ParentId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
