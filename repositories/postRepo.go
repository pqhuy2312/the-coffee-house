package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/models"
)

type PostRepo interface {
	Insert(p *models.Post) (*models.Post, error)
	Update(p *models.Post) (*models.Post, error)
	Delete(id int) (*models.Post, error)
	GetBySlug(slug string) (*models.Post, error)
	GetById(id int) (*models.Post, error)
}
