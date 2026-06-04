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
	if img.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file is over size(1 byte)",
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

func (n *NewsHandler) PostUploadFileNewsV1(ctx *gin.Context) {
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

	imageName, err := utils.ValidateAndSaveFile(img, "./new_file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"name":   input.Title,
		"Status": input.Status,
		"img":    imageName,
	})
}

func (n *NewsHandler) UploadMultipleFile(ctx *gin.Context) {
	var input dto.PostNewsV1
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	images := form.File["images"]

	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "no file was selected",
		})
		return
	}

	var (
		successFile []string
		failFile    []map[string]string
	)
	for _, file := range images {
		imageNames, err := utils.ValidateAndSaveFile(file, "./form-file")
		if err != nil {
			failFile = append(failFile, map[string]string{
				"file_name": file.Filename,
				"error":     err.Error(),
			})
			continue
		}

		successFile = append(successFile, imageNames)
	}

	resp := gin.H{
		"message":    "New file upload",
		"title":      input.Title,
		"status":     input.Status,
		"image_name": successFile,
	}

	if len(failFile) > 0 {
		resp["message"] = "Upload completed with partial error"
		resp["error"] = failFile
	}

	ctx.JSON(http.StatusOK, resp)
}
