package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/dto"
	"backend/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

type GetUserByUUIDParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

//Basic CRUD api

func (u *UserHandler) GetAllUser(ctx *gin.Context) {
	user, err := u.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidatorErrors(err))
		return
	}
	utils.ReponseSuccses(ctx, http.StatusAccepted, dto.MapUserToDTOs(user))
}
func (u *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	user, err := u.userService.FindByUUID(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(user)
	utils.ReponseSuccses(ctx, http.StatusOK, userDTO)

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
	userRequest := dto.MapUserToDTO(createUser)
	utils.ReponseSuccses(ctx, http.StatusCreated, &userRequest)
}
func (u *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (u *UserHandler) DeleteUser(ctx *gin.Context) {

}
