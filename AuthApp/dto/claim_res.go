package dto

import (
	"AuthApp/models"

	"github.com/golang-jwt/jwt/v4"
)

type IdTokenClaims struct {
	jwt.RegisteredClaims
	User *models.JWTuser `json:"user"`
}
