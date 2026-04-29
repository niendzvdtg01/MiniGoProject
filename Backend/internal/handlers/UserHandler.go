package handlers

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) GetUser(ctx *gin.Context) {
	
}
