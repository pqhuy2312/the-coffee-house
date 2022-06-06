package models

import (
	"time"
)


type Product struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique"`
	Slug string `json:"slug" gorm:"unique"`
	Info string `json:"info"`
	Price uint64  `json:"price"`
	Story string `json:"story"`
	Images []ProductImage `json:"images" gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;"`
	Sizes []Size `json:"sizes" gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;"`
	Toppings []Topping `json:"toppings" gorm:"many2many:product_toppings;constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
