package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	expectedKey := os.Getenv("API_KEY")

	if expectedKey == "" {
		expectedKey = "secret-key"
	}

	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("x-api-key")
		log.Println("x-api-key", apiKey)
		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "missing api key"})
			return
		}

		if apiKey != expectedKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":        "invalid api-key",
				"expected-key": expectedKey,
			})
			return
		}
		ctx.Set("username", "Nien")
		ctx.Next()
	}
}
