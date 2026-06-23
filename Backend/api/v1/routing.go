package api

import (
	"backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type Routing interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Routing) {
	r.Use(middlewares.LoggerMiddleware(), middlewares.AuthMiddleware())
	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}
