package service

import (
	"Backend/internal/model"

	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	FindUserByID(ctx *gin.Context) (*model.User, error)
}
