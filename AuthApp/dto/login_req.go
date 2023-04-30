package dto

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
