package main

import (
	"go-register/db"
	"go-register/server"
	"log"
)

func main() {
	err := db.Connect()

	if err != nil {
		log.Println("Failed to connect to DB")
	}

	server.Init()
}
