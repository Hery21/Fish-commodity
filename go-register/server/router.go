package server

import (
	"go-register/handlers"
	"go-register/services"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	RegisterService services.RegisterService
	AuthService     services.AuthService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.NewHandler(&handlers.HandlerConfig{
		RegisterService: c.RegisterService,
		AuthService:     c.AuthService,
	})

	router.POST("/register", h.RegisterHandler)

	router.POST("/login", h.LoginHandler)

	return router
}
