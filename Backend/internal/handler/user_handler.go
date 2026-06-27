package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/utils"
	"log"
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
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, utils.HandleValidatorErrors(err))
		return
	}

	createUser, err := u.userService.CreateUser(user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ReponseSuccses(ctx, http.StatusCreated, createUser)
}
func (u *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (u *UserHandler) DeleteUser(ctx *gin.Context) {

}
