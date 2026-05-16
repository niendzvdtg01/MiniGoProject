package main

import (
	"Backend/api"
	"Backend/pkg/utils"
)

func main() {
	server := api.SetupRouter()
	if err := utils.RegisterValidation(); err != nil {
		panic(err)
	}
	server.Run(":8085")
}
