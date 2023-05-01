package models

import "gorm.io/gorm"

type JWTuser struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Password   string `json:"-"`
}
