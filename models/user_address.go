package models

import (
	"time"
)


type UserAddress struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	UserId int `json:"userId"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
