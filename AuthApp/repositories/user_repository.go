package repositories

import (
	"errors"
	"go-register/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	MatchingCredential(phone string) (*models.JWTuser, error)
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

func (u *userRepository) MatchingCredential(phone string) (*models.JWTuser, error) {
	var user *models.JWTuser

	res := u.db.Table("users").Where("phone = ?", phone).First(&user)

	isNotFound := errors.Is(res.Error, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, res.Error
	}
	return user, nil
}
