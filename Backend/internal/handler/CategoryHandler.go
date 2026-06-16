package handler

import (
	"Backend/pkg/dto"
	"Backend/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) PostCategoryHandler(ctx *gin.Context) {
	var input dto.PostCategoryParam
	if err := ctx.ShouldBindQuery(&input); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.HandleValidatorErrors(err),
		})
		return
	}

	value, exists := ctx.Get("username")

	if !exists {
		log.Println(value)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalide username",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"name":     input.Name,
		"Status":   input.Status,
		"username": value,
	})

}
