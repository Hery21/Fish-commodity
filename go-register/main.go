package main

import (
	"go-register/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// err := db.Connect()

	// if err != nil {
	// 	log.Println("Failed to connect to DB")
	// }

	// server.Init()

	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)

	router.Run()
}

// type RegisterRes struct {
// 	Phone string
// 	Name string
// 	Role string
// 	Password int
// }
