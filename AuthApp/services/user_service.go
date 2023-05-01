package services

import (
	"AuthApp/config"
	"AuthApp/dto"
	"AuthApp/models"
	"AuthApp/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(*dto.LoginReq) (*dto.TokenResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
	appConfig      config.AppConfig
}

type AuthSConfig struct {
	UserRepository repositories.UserRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.JWTuser `json:"user"`
}

func (a *authService) generateJWTToken(user *models.JWTuser) (*dto.TokenResponse, error) {
	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}

	claims := &idTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			IssuedAt:  &timeNow,
			ExpiresAt: &timeExpire,
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.appConfig.JWTSecret)

	if err != nil {
		return new(dto.TokenResponse), err
	}
	return &dto.TokenResponse{IDToken: tokenString}, nil
}

func (a *authService) Login(req *dto.LoginReq) (*dto.TokenResponse, error) {
	user, err := a.userRepository.MatchingCredential(req.Phone)

	errNotMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if errNotMatch != nil || user == nil {
		return nil, err
	}

	token, err := a.generateJWTToken(user)

	return token, err
}
