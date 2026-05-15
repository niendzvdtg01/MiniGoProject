package handler

import (
	"Backend/pkg/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{
		"message": "Get user by id",
		"id":      id,
	})
}

func (u *UserHandler) GetUserByName(ctx *gin.Context) {
	var userRequest dto.UserRequest
	err := ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err: ": err,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"username": userRequest.UserName,
		"password": userRequest.Password,
	})
}
