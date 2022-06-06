package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagRepoImpl struct {
	Db *gorm.DB
}

func NewTagRepo(db *gorm.DB) repositories.TagRepo {
	return &TagRepoImpl {
		Db: db,
	}
}

func(t *TagRepoImpl) Insert(f *forms.FTag) (*models.Tag, error) {
	output := models.Tag{
		Title: f.Title,
		Slug: *f.Slug,
		TopicId: f.TopicId,
	}
	
	err := t.Db.Create(&output).Error

	return &output, err
}

func (t *TagRepoImpl) GetBySlug(slug string) (*models.Tag, error) {
	var output models.Tag
	
	err := t.Db.Where("slug = ?", slug).First(&output).Error
	
	return &output, err
}

func (t *TagRepoImpl) GetById(id int) (*models.Tag, error) {
	var output models.Tag
	
	err := t.Db.First(&output, id).Error
	
	return &output, err
}

func (c *TagRepoImpl) Delete(id int) (*models.Tag, error) {
	var output *models.Tag

	err := c.Db.Clauses(clause.Returning{}).Delete(&output, id).Error

	return output, err
}

func (c *TagRepoImpl) Update(fTag *models.Tag, id int) (*models.Tag , error) {
	output := models.Tag{}

	err := c.Db.Model(&output).Clauses(clause.Returning{}).Where("id = ?", id).Updates(fTag).Error

	if err != nil {
		return nil, err
	}

	return &output, nil
}