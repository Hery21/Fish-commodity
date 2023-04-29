package handlers

import (
	"go-register/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var registerReq dto.RegisterReq

	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Phone": registerReq.Phone,
		"Name":  registerReq.Name,
	})
}
