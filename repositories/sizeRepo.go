package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type SizeRepo interface {
	Insert(u *forms.FSize) (*models.Size, error)
	GetByName(name string) (*models.Size, error)
}
