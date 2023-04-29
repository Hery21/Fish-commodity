package server

import (
	"go-register/handlers"
	"go-register/services"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	RegisterService services.RegisterService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handlers.NewHandler(&handlers.HandlerConfig{
		RegisterService: c.RegisterService,
	})

	router.POST("/register", h.RegisterHandler)

	return router
}
