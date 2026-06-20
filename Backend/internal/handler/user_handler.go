package handler

import (
	"backend/internal/service"
	"log"

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
	u.userService.FindAll()
	log.Println("dcmmm")
}
func (u *UserHandler) GetUserByUUID(ctx *gin.Context) {

}
func (u *UserHandler) CreateUser(ctx *gin.Context) {

}
func (u *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (u *UserHandler) DeleteUser(ctx *gin.Context) {

}
