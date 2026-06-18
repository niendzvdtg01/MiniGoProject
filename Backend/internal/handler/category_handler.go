package handler

import (
	"backend/pkg/dto"
	"backend/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var request dto.CategoryRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.HandleValidatorErrors(err),
		})
		return
	}

	value, exists := ctx.Get("username")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid username",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"name":     request.Name,
		"status":   request.Status,
		"username": value,
	})

}
