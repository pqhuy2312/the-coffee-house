package models

import (
	"time"
)


type Size struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	ProductId int    `json:"productId"`
	Price uint64  `json:"price"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
