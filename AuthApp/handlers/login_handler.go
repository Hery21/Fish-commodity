package handlers

import (
	"go-register/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	var value dto.LoginReq
	err := c.ShouldBindJSON(&value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	input, err := h.authService.Login(&value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, input)
}
