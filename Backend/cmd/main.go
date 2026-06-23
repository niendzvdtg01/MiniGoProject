package main

import (
	"backend/internal/app"
	"backend/internal/config"
	"backend/internal/middlewares"
	"backend/pkg/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	go middlewares.CleanupClients()

	utils.InitValidation()
	//Init config
	cfg := config.NewConfig()
	//Init application
	application := app.NewApplication(cfg)
	//start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}
