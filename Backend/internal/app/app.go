package app

import (
	"backend/api/v1"
	"backend/internal/config"

	"github.com/gin-gonic/gin"
)

type Module interface {
	Routes() api.Routing
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()
	modules := []Module{
		NewUserModule(),
	}
	api.RegisterRoutes(r, getModuleRoute(modules)...)
	return &Application{
		router: r,
		config: cfg,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServeAddress)
}

func getModuleRoute(modules []Module) []api.Routing {
	routeList := make([]api.Routing, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
