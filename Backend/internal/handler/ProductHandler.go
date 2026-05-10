package handler

import (
	"Backend/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

//validate data

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) PostProducts(ctx *gin.Context) {
	var input dto.PostProduct

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"body": string(body),
		"name": input.Name,
	})
}
