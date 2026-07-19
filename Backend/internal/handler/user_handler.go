package handler

import (
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

type GetuserParam struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50"`
	Page   int    `form:"page" binding:"omitempty,gte=1,lte=100"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

//Basic CRUD api

func (u *UserHandler) GetAllUser(ctx *gin.Context) {
	var params GetuserParam

	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, utils.HandleValidatorErrors(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	log.Println(params)

	user, err := u.userService.FindAll(params.Search, params.Page, params.Limit)
	if err != nil {
		utils.ResponseValidator(ctx, utils.HandleValidatorErrors(err))
		return
	}
	userDto := dto.MapUserToDTOs(user)
	log.Println(userDto)
	utils.ReponseSuccses(ctx, http.StatusAccepted, userDto)
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
	var input dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		utils.ResponseValidator(ctx, utils.HandleValidatorErrors(err))
		return
	}

	// for ToModel
	user := input.MapCreateInputToModel()
	createUser, err := u.userService.CreateUser(user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	userRequest := dto.MapUserToDTO(createUser)
	utils.ReponseSuccses(ctx, http.StatusCreated, &userRequest)
}
func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, utils.HandleValidatorErrors(err))
		return
	}

	var input dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		utils.ResponseValidator(ctx, utils.HandleValidatorErrors(err))
		return
	}

	user := input.MapUpdateInputToModel()

	updateUser, err := u.userService.UpdateUser(params.Uuid, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(updateUser)
	log.Println(userDTO)
	utils.ReponseSuccses(ctx, http.StatusOK, &userDTO)

}
func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, utils.HandleValidatorErrors(err))
		return
	}

	if err := u.userService.DeleteUser(params.Uuid); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ReponseStatusCode(ctx, http.StatusNoContent)
}
