package handler

import (
	"Backend/pkg/dto"
	"Backend/pkg/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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
	img, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file is a required",
		})
		return
	}
	//have permission
	err = os.MkdirAll("./uploads", os.ModePerm)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create folder!"})
		return
	}

	dst := fmt.Sprintf("./uploads/%s", filepath.Base(img.Filename))

	if err := ctx.SaveUploadedFile(img, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not save file"})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"name":   input.Title,
		"Status": input.Status,
		"img":    img.Filename,
		"path":   dst,
	})
}
