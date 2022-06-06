package repositories

import (
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type UserRepo interface {
	Insert(u *forms.FRegister) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	Update(u *models.User) (*models.User, error)
}
