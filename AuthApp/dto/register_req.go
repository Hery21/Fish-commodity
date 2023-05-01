package dto

type RegisterReq struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Role  string `json:"role" binding:"required"`
}
