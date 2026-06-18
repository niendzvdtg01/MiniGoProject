package api

import (
	"backend/internal/handler"
	"backend/internal/middlewares"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(productService *service.ProductService) *gin.Engine {
	router := gin.Default()
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler(productService)
	categoryHandler := handler.NewCategoryHandler()
	newsHandler := handler.NewNewsHandler()
	messageHandler := handler.NewMessageHandler()
	router.Use(middlewares.LoggerMiddleware())
	apiGroup := router.Group("/api")
	{
		userGroup := apiGroup.Group("/user").Use(middlewares.RateLimitingMiddleware())
		{
			userGroup.GET("/:id", middlewares.ApiKeyMiddleware(), userHandler.GetUserByID)
			userGroup.POST("/info", userHandler.GetUserByUsername)
		}
		productGroup := apiGroup.Group("/product").Use(middlewares.LoggerMiddleware())
		{
			productGroup.POST("/all", productHandler.CreateProduct)
			productGroup.POST("/product_info", productHandler.CreateProductFromForm)
		}
		categoryGroup := apiGroup.Group("/category").Use(middlewares.ApiKeyMiddleware())
		{
			categoryGroup.POST("", categoryHandler.CreateCategory)
		}
		newsGroup := apiGroup.Group("/news")
		{
			newsGroup.POST("", newsHandler.CreateNews)
			newsGroup.POST("/image", newsHandler.UploadNewsImage)
			newsGroup.POST("/upload-multiple-file", newsHandler.UploadMultipleFiles)
		}
		messageGroup := apiGroup.Group("/message")
		{
			messageGroup.POST("/post_message", messageHandler.PostMessage)
		}
	}

	router.StaticFS("/images", gin.Dir("./form-file", false))
	return router
}
