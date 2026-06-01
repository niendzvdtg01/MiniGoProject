package api

import (
	"Backend/internal/handler"
	"Backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(productService *service.ProductService) *gin.Engine {
	server := gin.Default()
	//define user handler

	userHandler := handler.NewUserHandler()
	//product handler
	productHandler := handler.NewProductHandler(productService)
	categoryHandler := handler.NewCategoryHandler()
	newsHandler := handler.NewNewsHandler()
	messageHandler := handler.NewMessageHandler()
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
		categoryAPI := serverRouting.Group("/category")
		{
			categoryAPI.POST("", categoryHandler.PostCategoryHandler)
		}
		newsAPI := serverRouting.Group("/news")
		{
			newsAPI.POST("", newsHandler.PostNewsV1)
		}
		messageAPI := serverRouting.Group("/message")
		{
			messageAPI.POST("/post_message", messageHandler.PostMessgaeV1)
		}
	}
	return server
}
