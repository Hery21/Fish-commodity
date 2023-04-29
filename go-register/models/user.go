package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Password   string `json:"password"`
}
