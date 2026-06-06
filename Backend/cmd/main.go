package main

import (
	"Backend/api/v1"
	"Backend/internal/service"
	"Backend/pkg/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//service declare

	if err := godotenv.Load(); err != nil {
		log.Println("No .env variables were found")
	}
	productService := service.NewProductService()
	server := api.SetupRouter(productService)
	if err := utils.RegisterValidation(); err != nil {
		panic(err)
	}
	server.Run(":8085")
}
