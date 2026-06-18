package handler

import (
	"backend/pkg/dto"
	"backend/pkg/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct{}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (h *NewsHandler) CreateNews(ctx *gin.Context) {
	var request dto.NewsFormRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.HandleValidatorErrors(err),
		})
		return
	}
	img, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	if img.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file exceeds the 5 MB limit",
		})
		return
	}

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create folder!"})
		return
	}

	dst := fmt.Sprintf("./uploads/%s", filepath.Base(img.Filename))

	if err := ctx.SaveUploadedFile(img, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not save file"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"name":   request.Title,
		"status": request.Status,
		"img":    img.Filename,
		"path":   dst,
	})
}

func (h *NewsHandler) UploadNewsImage(ctx *gin.Context) {
	var request dto.NewsFormRequest
	if err := ctx.ShouldBind(&request); err != nil {
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
		"name":   request.Title,
		"status": request.Status,
		"img":    imageName,
	})
}

func (h *NewsHandler) UploadMultipleFiles(ctx *gin.Context) {
	const publicUrl = "http://localhost:8085/images/"

	var request dto.NewsFormRequest
	if err := ctx.ShouldBind(&request); err != nil {
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
		publicImageUrl := publicUrl + imageNames

		successFile = append(successFile, publicImageUrl)
	}
	//public images  ye ye hu hu

	resp := gin.H{
		"message":    "New file upload",
		"title":      request.Title,
		"status":     request.Status,
		"image_name": successFile,
	}

	if len(failFile) > 0 {
		resp["message"] = "Upload completed with partial error"
		resp["error"] = failFile
	}

	ctx.JSON(http.StatusOK, resp)
}
