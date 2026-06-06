package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Start func - Check from middleware"))
		//before handler
		ctx.Next()
		log.Println("End")
	}
}
