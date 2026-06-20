package app

import (
	"backend/api/v1"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
)

type UserModule struct {
	routes api.Routing
}

func NewUserModule() *UserModule {
	//Initialize repository
	userRepo := repository.NewUserRepository()
	//Initialize service
	userService := service.NewUserService(userRepo)
	//Inittializr handler
	userHandler := handler.NewUserHandler(userService)
	//Initialize routes
	userRoutes := api.NewUserRoutes(userHandler)
	return &UserModule{routes: userRoutes}
}

func (m *UserModule) Routes() api.Routing {
	return m.routes
}
