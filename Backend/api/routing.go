package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	server := gin.Default()

	serverRouting := server.Group("/api")
	{
		userApi := serverRouting.Group("/user")
		{
			userApi.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusAccepted, gin.H{"message": "Successfully!!"}) })
		}
	}
	return server
}
