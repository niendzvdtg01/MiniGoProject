package api

import (
	"Backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	server := gin.Default()
	//define user handler

	userHandler := handler.NewUserHandler()
	//product handler
	productHandler := handler.NewProductHandler()
	serverRouting := server.Group("/api")
	{
		userApi := serverRouting.Group("/user")
		{
			userApi.GET("/", userHandler.GetUserByID)
		}
		productAPI := serverRouting.Group("/product")
		{
			productAPI.POST("", productHandler.PostProducts)
		}
	}
	return server
}
