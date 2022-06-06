package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type CategoryRepo interface {
	List() ([]*models.Category, error)
	Insert(c *forms.FCategory) (*models.Category, error)
	GetCategoryBySlug(slug string) (*models.Category, error)
	GetCategoryById(id int) (*models.Category, error)
	Delete(id int) (*models.Category, error)
	Update(c *forms.FCategory, id int) (*models.Category, error)
}