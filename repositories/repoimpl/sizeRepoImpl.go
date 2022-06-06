package repoimpl

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"gorm.io/gorm"
)

type sizeImpl struct {
	Db *gorm.DB
}

func NewSizeRepo(db *gorm.DB) repositories.SizeRepo {
	return &sizeImpl {
		Db: db,
	}
}

func (ps *sizeImpl) Insert(fPs *forms.FSize) (*models.Size, error) {
	output := models.Size{
		Name: fPs.Name,
	}

	err := ps.Db.Create(&output).Error

	return &output, err
}

func (ps *sizeImpl) GetByName(name string) (*models.Size, error) {
	var output models.Size

	err := ps.Db.Where("name = ?", name).First(&output).Error

	return &output, err
}