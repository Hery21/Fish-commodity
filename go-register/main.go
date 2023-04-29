package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// err := db.Connect()

	// if err != nil {
	// 	log.Println("Failed to connect to DB")
	// }

	// server.Init()

	router := gin.Default()

	router.POST("/register", RegisterHandler)

	router.Run()
}

type RegisterReq struct {
	Phone string
	Name  string
	Role  string
}

// type RegisterRes struct {
// 	Phone string
// 	Name string
// 	Role string
// 	Password int
// }

func RegisterHandler(c *gin.Context) {
	var registerReq RegisterReq

	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Phone": registerReq.Phone,
		"Name":  registerReq.Name,
	})
}
