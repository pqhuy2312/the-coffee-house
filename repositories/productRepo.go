package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/models"
)

type ProductRepo interface {
	Insert(p *models.Product) (*models.Product, error)
	Update(p *models.Product) (*models.Product, error)
	Delete(id int) (*models.Product, error)
	GetBySlug(slug string) (*models.Product, error)
	GetById(id int) (*models.Product, error)
}
