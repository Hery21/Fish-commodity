package handlers

import "go-register/services"

type HandlerConfig struct {
	RegisterService services.RegisterService
}

type Handler struct {
	registerService services.RegisterService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		registerService: c.RegisterService,
	}
}
