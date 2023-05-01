package repositories

import (
	"AuthApp/models"
	"fmt"

	"gorm.io/gorm"
)

type RegisterRepository interface {
	Register(user *models.User) (*models.User, error)
}

type registerRepository struct {
	db *gorm.DB
}

type RRConfig struct {
	DB *gorm.DB
}

func NewRegisterRepository(rr *RRConfig) *registerRepository {
	return &registerRepository{db: rr.DB}
}

func (r *registerRepository) Register(user *models.User) (*models.User, error) {
	var existingUser models.User
	if err := r.db.Where("phone = ?", user.Phone).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
