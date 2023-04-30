package handlers

import (
	"go-register/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	// value, _ := c.ShouldBindJSON(&model)
	// fmt.Println("lewat====================", value)
	var value dto.RegisterReq
	err := c.ShouldBindJSON(&value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	input, err := h.registerService.Register(&value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, input)
}
