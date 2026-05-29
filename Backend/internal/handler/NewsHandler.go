package handler

import (
	"Backend/pkg/dto"
	"Backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) PostNewsV1(ctx *gin.Context) {
	var input dto.PostNewsV1
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.HandleValidatorErrors(err),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"name":   input.Title,
		"Status": input.Status,
	})
}
