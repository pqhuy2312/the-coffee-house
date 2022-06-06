package models

import (
	"time"
)


type Topping struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique"`
	Price uint64  `json:"price"` 
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
