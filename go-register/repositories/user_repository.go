package repositories

import (
	"errors"
	"fmt"
	"go-register/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	MatchingCredential(phone string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) *userRepository {
	return &userRepository{db: c.DB}
}

func (u *userRepository) MatchingCredential(phone string) (*models.User, error) {
	var user *models.User
	fmt.Println("lewat================")
	fmt.Println(phone)

	res := u.db.Where("phone = ?", phone).First(&user)

	isNotFound := errors.Is(res.Error, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, res.Error
	}
	return user, nil
}
