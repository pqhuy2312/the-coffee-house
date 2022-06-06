package repoimpl

import (
	"log"

	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories"
	"github.com/pqhuy2312/the-coffee-house/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &UserRepoImpl {
		Db: db,
	}
}

func (u *UserRepoImpl) Insert(user *forms.FRegister) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		UserName: user.UserName,
		Email: user.Email,
		Password: hashedPassword,
	}

	err = u.Db.Create(&newUser).Error

	return &newUser, err
}

func (u *UserRepoImpl) GetUserByEmail(email string) (*models.User, error) {
	var output models.User
	
	err := u.Db.Where("email = ?", email).First(&output).Error
	
	return &output, err
}

func (u *UserRepoImpl) GetUserById(id int) (*models.User, error) {
	var output models.User
	
	err := u.Db.First(&output, id).Error
	
	return &output, err
}

func (u *UserRepoImpl) Update(user *models.User) (*models.User, error) {
	output := models.User{}
	err := u.Db.Model(&output).Clauses(clause.Returning{}).Where("id = ?", user.Id).Updates(user).Error

	if err != nil {
		return nil, err
	}

	return &output, nil
}