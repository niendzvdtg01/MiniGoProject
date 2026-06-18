package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Start func - Check from middleware"))
		ctx.Next()
		log.Println("End")
	}
}
