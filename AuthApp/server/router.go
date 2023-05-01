package server

import (
	"AuthApp/handlers"
	"AuthApp/middlewares"
	"AuthApp/services"

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

	router.GET("/validate", middlewares.AuthorizeJWT, h.ValidateHandler)

	return router
}
