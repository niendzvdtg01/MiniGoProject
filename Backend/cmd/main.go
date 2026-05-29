package main

import (
	"Backend/api/v1"
	"Backend/internal/service"
	"Backend/pkg/utils"
)

func main() {
	//service declare
	productService := service.NewProductService()
	server := api.SetupRouter(productService)
	if err := utils.RegisterValidation(); err != nil {
		panic(err)
	}
	server.Run(":8085")
}
