package main

import (
	"backend/internal/app"
	"backend/internal/config"
	"backend/internal/db"
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

	if err := utils.InitValidation(); err != nil {
		log.Fatalf("Validator init fail!: %v", err)
	}
	if err := db.InitDB(); err != nil {
		log.Fatal("Fail to connect db", err)
	}
	//Init config
	cfg := config.NewConfig()
	//Init application
	application := app.NewApplication(cfg)
	//start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}
