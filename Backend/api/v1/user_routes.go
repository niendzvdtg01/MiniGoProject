package api

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{handler: handler}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/user")
	{
		users.GET("", ur.handler.GetAllUser)
		users.GET("/:uuid", ur.handler.GetUserByUUID)
		users.POST("", ur.handler.CreateUser)
		users.PUT("/:uuid", ur.handler.UpdateUser)
		users.DELETE("/:uuid", ur.handler.DeleteUser)
	}
}
