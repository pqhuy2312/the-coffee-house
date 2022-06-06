package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type TopicRepo interface {
	Insert(f *forms.FTopic) (*models.Topic, error)
	GetBySlug(slug string) (*models.Topic, error)
	GetById(id int) (*models.Topic, error)
	Delete(id int) (*models.Topic, error)
	Update(c *forms.FTopic, id int) (*models.Topic, error)
}