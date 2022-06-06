package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productImpl struct {
	Db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repositories.ProductRepo {
	return &productImpl {
		Db: db,
	}
}

func (p *productImpl) Insert(f *models.Product) (*models.Product, error) {
	var output = f
	tx := p.Db.Begin()

	err := tx.Create(&output).Error

	if err != nil {
		return nil, err
	}

	return output, tx.Commit().Error
}

func (p *productImpl) Update(input *models.Product) (*models.Product, error) {
	output := input
	tx := p.Db.Begin()
	err := tx.Where("product_id = ?", output.Id).Delete(&models.Size{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Where("product_id = ?", output.Id).Delete(&models.ProductImage{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Model(&output).Association("Toppings").Replace(output.Toppings)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&output).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return output, tx.Commit().Error
}


func (p *productImpl) GetBySlug(slug string) (*models.Product, error) {
	var output *models.Product

	err := p.Db.Where("slug = ?", slug).First(&output).Error
	
	return output, err
}

func (p *productImpl) GetById(id int) (*models.Product, error) {
	var output *models.Product

	err := p.Db.Preload("Sizes").Preload("Toppings").First(&output, id).Error
	
	return output, err
}

func (p *productImpl) Delete(id int) (*models.Product, error) {
	var output *models.Product

	err := p.Db.Clauses(clause.Returning{}).Delete(&output, id).Error
	
	return output, err
}