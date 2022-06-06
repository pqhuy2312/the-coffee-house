package models

import "time"

type User struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	UserName     string    `json:"userName"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Avatar       string    `json:"avatar"`
	Role         string    `json:"role" gorm:"default:'USER'"`
	SocialId     *string   `json:"socialId"`
	Address      *UserAddress    `json:"address" gorm:"foreignKey:UserId"`
	TokenVersion int       `json:"tokenVersion"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
