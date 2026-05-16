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
			userApi.GET("/:id", userHandler.GetUserByID)
			userApi.POST("/info", userHandler.GetUserByName)
		}
		productAPI := serverRouting.Group("/product")
		{
			productAPI.POST("/all", productHandler.PostProducts)
			productAPI.POST("/product_info", productHandler.PostProductRequest)
		}
	}
	return server
}
