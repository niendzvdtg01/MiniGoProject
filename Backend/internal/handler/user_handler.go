package handler

import (
	"backend/internal/service"
	"backend/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by id",
		"id":      id,
	})
}

func (h *UserHandler) GetUserByUsername(ctx *gin.Context) {
	var request dto.UserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"username": request.UserName,
		"password": request.Password,
	})
}
