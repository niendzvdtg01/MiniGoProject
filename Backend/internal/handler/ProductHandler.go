package handler

import (
	"Backend/pkg/dto"
	"Backend/pkg/utils"
	"fmt"
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
	input.Display = true

	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":          input.Name,
		"product_image": input.ProductImage,
		"display":       input.Display,
	})
}

func (p *ProductHandler) PostProductRequest(ctx *gin.Context) {
	var input dto.ProductRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": utils.HandleValidatorErrors(err)})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"product_name": input.ProductName,
		"best_saling":  input.BestSaling,
		"price":        input.Price,
	})
}
