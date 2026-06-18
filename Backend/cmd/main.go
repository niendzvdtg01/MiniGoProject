package main

import (
	"backend/api/v1"
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/middlewares"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	go middlewares.CleanupClients()

	if err := utils.RegisterValidation(); err != nil {
		panic(err)
	}
	cfg := config.NewConfig()

	//Initialize repository
	userRepo := repository.NewUserRepository()
	//Initialize service
	userService := service.NewUserService(userRepo)
	//Inittializr handler
	userHandler := handler.NewUserHandler(userService)
	//Initialize routes
	userRoutes := api.NewUserRoutes(userHandler)

	//setup server
	r := gin.Default()

	api.RegisterRoutes(r, userRoutes)

	if err := r.Run(cfg.ServeAddress); err != nil {
		panic(err)
	}
}
