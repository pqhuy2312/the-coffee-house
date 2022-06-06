package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepoImpl struct {
	Db *gorm.DB
}

func NewPostRepo(db *gorm.DB) repositories.PostRepo {
	return &PostRepoImpl {
		Db: db,
	}
}

func(p *PostRepoImpl) Insert(post *models.Post) (*models.Post, error) {
	output := post
	
	err := p.Db.Create(&output).Error

	return output, err
}

func (p *PostRepoImpl) GetBySlug(slug string) (*models.Post, error) {
	var output models.Post
	
	err := p.Db.Preload("Tag").Preload("User").Where("slug = ?", slug).First(&output).Error
	
	return &output, err
}

func (p *PostRepoImpl) GetById(id int) (*models.Post, error) {
	var output models.Post
	
	err := p.Db.Preload("Tag").Preload("Author").First(&output, id).Error
	
	return &output, err
}

func (p *PostRepoImpl) Delete(id int) (*models.Post, error) {
	var output *models.Post

	err := p.Db.Clauses(clause.Returning{}).Delete(&output, id).Error

	return output, err
}

func (p *PostRepoImpl) Update(fPost *models.Post) (*models.Post , error) {
	output := fPost

	err := p.Db.Model(&output).Clauses(clause.Returning{}).Where("id = ?", fPost.Id).Updates(fPost).Error

	if err != nil {
		return nil, err
	}

	return output, nil
}