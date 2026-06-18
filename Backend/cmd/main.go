package main

import (
	"backend/api/v1"
	"backend/internal/middlewares"
	"backend/internal/service"
	"backend/pkg/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	productService := service.NewProductService()
	router := api.SetupRouter(productService)

	go middlewares.CleanupClients()

	if err := utils.RegisterValidation(); err != nil {
		panic(err)
	}
	if err := router.Run(":8085"); err != nil {
		log.Fatal(err)
	}
}
