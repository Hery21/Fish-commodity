package dto

import (
	"AuthApp/models"
)

type RegisterRes struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (rr *RegisterRes) FromRegister(r *models.User) *RegisterRes {
	return &RegisterRes{
		Phone:    r.Phone,
		Name:     r.Name,
		Role:     r.Role,
		Password: r.Password,
	}
}
