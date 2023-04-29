package server

import (
	"fmt"
	"go-register/db"
	"go-register/repositories"
	"go-register/services"
)

func Init() {
	registerRepository := repositories.NewRegisterRepository(&repositories.RRConfig{DB: db.Get()})
	registerService := services.NewRegisterService(&services.RSConfig{RegisterRepository: registerRepository})

	router := NewRouter(&RouterConfig{
		RegisterService: registerService,
	})

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("server Error: ", err)
	}
}
