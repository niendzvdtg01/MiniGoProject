package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (m *MessageHandler) PostMessgaeV1(ctx *gin.Context) {
	msg, err := ctx.FormFile("message")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not upload message"})
		return
	}

	err = os.MkdirAll("./message", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not upload file"})
		return
	}

	dst := fmt.Sprintf("./message/%s", filepath.Base(msg.Filename))

	if err := ctx.SaveUploadedFile(msg, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can not save file"})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message":   "Successfully",
		"file_name": msg.Filename,
	})
}
