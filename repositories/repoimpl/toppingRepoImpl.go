package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"github.com/pqhuy2312/the-coffee-house/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type toppingImpl struct {
	Db *gorm.DB
}

func NewToppingRepo(db *gorm.DB) repositories.ToppingRepo {
	return &toppingImpl {
		Db: db,
	}
}

func (t *toppingImpl) Insert(f *forms.FTopping) (*models.Topping, error) {
	output := models.Topping{
		Name: f.Name,
		Price: f.Price,
	}

	err := t.Db.Create(&output).Error

	return &output, err
}

func (t *toppingImpl) GetByName(name string) (*models.Topping, error) {
	var output models.Topping

	err := t.Db.Where("name = ?", name).First(&output).Error

	return &output, err
}

func (t *toppingImpl) GetByIds(ids []int) ([]models.Topping, error) {
	var output []models.Topping

	err := t.Db.Where(ids).Find(&output).Error

	return output, err
}

func (t *toppingImpl) GetById(id int) (models.Topping, error) {
	var output models.Topping

	err := t.Db.First(&output, id).Error

	return output, err
}

func (t *toppingImpl) Update(ft *forms.FTopping, id int) (*models.Topping, error) {
	output := models.Topping{}

	mapData := utils.StructToMap(ft)

	err := t.Db.Model(&output).Clauses(clause.Returning{}).Where("id = ?", id).Updates(mapData).Error

	return &output, err
}

func (t *toppingImpl) Delete(id int) (*models.Topping, error) {
	var output *models.Topping

	err := t.Db.Clauses(clause.Returning{}).Delete(&output, id).Error

	return output, err
}