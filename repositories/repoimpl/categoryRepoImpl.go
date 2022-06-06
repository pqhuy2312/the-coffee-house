package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"github.com/pqhuy2312/the-coffee-house/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepoImpl struct {
	Db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) repositories.CategoryRepo {
	return &CategoryRepoImpl {
		Db: db,
	}
}

func(c *CategoryRepoImpl) Insert(fcategory *forms.FCategory) (*models.Category, error) {
	output := models.Category{
		Title: *fcategory.Title,
		Slug: *fcategory.Slug,
		ParentId: fcategory.ParentId,
	}
	
	err := c.Db.Create(&output).Error

	return &output, err
}

func (c *CategoryRepoImpl) GetCategoryBySlug(slug string) (*models.Category, error) {
	var output models.Category
	
	err := c.Db.Where("slug = ?", slug).First(&output).Error
	
	return &output, err
}

func (c *CategoryRepoImpl) GetCategoryById(id int) (*models.Category, error) {
	var output models.Category
	
	err := c.Db.First(&output, id).Error
	
	return &output, err
}

func (c *CategoryRepoImpl) List() ([]*models.Category, error) {
	
	var output []*models.Category

	err := c.Db.Preload("Children").Where("parent_id IS NULL").Find(&output).Error

	if err != nil {
		return []*models.Category{}, err
	}
	
	return output, err
}

func (c *CategoryRepoImpl) Delete(id int) (*models.Category, error) {
	var output *models.Category

	err := c.Db.Clauses(clause.Returning{}).Delete(&output, id).Error

	return output, err
}

func (c *CategoryRepoImpl) Update(fCategory *forms.FCategory, id int) (*models.Category , error) {
	output := models.Category{}

	mapData := utils.StructToMap(fCategory)


	err := c.Db.Model(&output).Clauses(clause.Returning{}).Where("id = ?", id).Updates(mapData).Error

	if err != nil {
		return nil, err
	}

	return &output, nil
}
