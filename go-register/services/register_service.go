package services

import (
	"fmt"
	"go-register/dto"
	"go-register/models"
	"go-register/repositories"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type RegisterService interface {
	Register(req *dto.RegisterReq) (*dto.RegisterRes, error)
}

type registerService struct {
	registerRepository repositories.RegisterRepository
}

type RSConfig struct {
	RegisterRepository repositories.RegisterRepository
}

func NewRegisterService(r *RSConfig) RegisterService {
	return &registerService{
		registerRepository: r.RegisterRepository,
	}
}

func (r *registerService) Register(req *dto.RegisterReq) (*dto.RegisterRes, error) {
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, 4)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(string(password))

	bytes, _ := bcrypt.GenerateFromPassword([]byte(string(password)), 10)
	decryptedPassword := string(bytes)
	registeringUser := &models.User{
		Phone:    req.Phone,
		Name:     req.Name,
		Role:     req.Role,
		Password: decryptedPassword,
	}

	registeredUser, err := r.registerRepository.Register(registeringUser)

	if err != nil {
		return new(dto.RegisterRes), err
	}
	registeredUser.Password = string(password)

	return new(dto.RegisterRes).FromRegister(registeredUser), err
}
