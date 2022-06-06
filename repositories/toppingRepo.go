package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type ToppingRepo interface {
	Insert(t *forms.FTopping) (*models.Topping, error)
	Update(t *forms.FTopping, id int) (*models.Topping, error)
	Delete(id int) (*models.Topping, error)
	GetByName(name string) (*models.Topping, error)
	GetByIds(ids []int) ([]models.Topping, error)
	GetById(id int) (models.Topping, error)
}
