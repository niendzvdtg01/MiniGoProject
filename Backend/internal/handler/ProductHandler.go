package handler

import (
	"Backend/internal/service"
	"Backend/pkg/dto"
	"Backend/pkg/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

//validate data

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (p *ProductHandler) PostProducts(ctx *gin.Context) {
	var input dto.PostProduct
	input.Display = true

	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": utils.HandleValidatorErrors(err)})
		return
	}

	if err := p.productService.ValidateUUID(input); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":              input.Name,
		"product_image":     input.ProductImage,
		"display":           input.Display,
		"product_attribute": input.ProductAttribute,
		"product_info":      input.ProductInfo,
		"product_metadata":  input.ProductMetaData,
	})
}

func (p *ProductHandler) PostProductRequest(ctx *gin.Context) {
	var input dto.ProductRequest
	if err := ctx.ShouldBind(&input); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": utils.HandleValidatorErrors(err)})
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "file noot found!",
		})
	}

	fileName, err := utils.ValidateAndSaveFile(file, "./file")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"product_name": input.ProductName,
		"best_saling":  input.BestSaling,
		"price":        input.Price,
		"file_name":    fileName,
	})
}

func (p *ProductHandler) PostCategoryV1() {

}
