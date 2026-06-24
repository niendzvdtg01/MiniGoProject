package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

//Basic CRUD api

func (u *UserHandler) GetAllUser(ctx *gin.Context) {

}
func (u *UserHandler) GetUserByUUID(ctx *gin.Context) {

}
func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidatorErrors(err))
		return
	}
	u.userService.CreateUser(user)
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Succesful",
		"data":    user,
	})

}
func (u *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (u *UserHandler) DeleteUser(ctx *gin.Context) {

}
