package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"github.com/pqhuy2312/the-coffee-house/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TopicRepoImpl struct {
	Db *gorm.DB
}

func NewTopicRepo(db *gorm.DB) repositories.TopicRepo {
	return &TopicRepoImpl {
		Db: db,
	}
}

func(t *TopicRepoImpl) Insert(f *forms.FTopic) (*models.Topic, error) {
	output := models.Topic{
		Title: f.Title,
		Slug: *f.Slug,
	}
	
	err := t.Db.Create(&output).Error

	return &output, err
}

func (t *TopicRepoImpl) GetBySlug(slug string) (*models.Topic, error) {
	var output models.Topic
	
	err := t.Db.Where("slug = ?", slug).First(&output).Error
	
	return &output, err
}

func (t *TopicRepoImpl) GetById(id int) (*models.Topic, error) {
	var output models.Topic
	
	err := t.Db.First(&output, id).Error
	
	return &output, err
}

func (c *TopicRepoImpl) Delete(id int) (*models.Topic, error) {
	var output *models.Topic

	err := c.Db.Clauses(clause.Returning{}).Delete(&output, id).Error

	return output, err
}

func (c *TopicRepoImpl) Update(fTopic *forms.FTopic, id int) (*models.Topic , error) {
	output := models.Topic{}

	mapData := utils.StructToMap(fTopic)


	err := c.Db.Model(&output).Clauses(clause.Returning{}).Where("id = ?", id).Updates(mapData).Error

	if err != nil {
		return nil, err
	}

	return &output, nil
}