package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type TagRepo interface {
	Insert(f *forms.FTag) (*models.Tag, error)
	GetBySlug(slug string) (*models.Tag, error)
	GetById(id int) (*models.Tag, error)
	Delete(id int) (*models.Tag, error)
	Update(c *models.Tag, id int) (*models.Tag, error)
}