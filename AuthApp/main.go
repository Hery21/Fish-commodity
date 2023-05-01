package main

import (
	"AuthApp/db"
	"AuthApp/server"
	"log"
)

func main() {
	err := db.Connect()

	if err != nil {
		log.Println("Failed to connect to DB")
	}

	server.Init()
}
