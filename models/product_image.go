package models

import (
	"time"
)


type ProductImage struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Url     string `json:"url"`
	ProductId     string `json:"productId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
