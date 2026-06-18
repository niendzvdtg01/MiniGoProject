package handler

import (
	"backend/internal/service"
	"backend/pkg/dto"
	"backend/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var request dto.CreateProductRequest
	request.Display = true

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.HandleValidatorErrors(err)})
		return
	}

	if err := h.productService.ValidateProductInfoKeys(request); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":              request.Name,
		"product_image":     request.ProductImage,
		"display":           request.Display,
		"product_attribute": request.ProductAttribute,
		"product_info":      request.ProductInfo,
		"product_metadata":  request.ProductMetaData,
	})
}

func (h *ProductHandler) CreateProductFromForm(ctx *gin.Context) {
	var request dto.ProductFormRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.HandleValidatorErrors(err)})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	fileName, err := utils.ValidateAndSaveFile(file, "./file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"product_name": request.ProductName,
		"best_saling":  request.BestSaling,
		"price":        request.Price,
		"file_name":    fileName,
	})
}
