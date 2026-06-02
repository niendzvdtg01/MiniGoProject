package utils

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

var allowExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

const maxSize = 5 << 20

// ValidateAndUploadFile(file, "./uploads")
func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	ext := strings.ToLower(filepath.Ext((fileHeader.Filename)))

	if !allowExts[ext] {
		return "", errors.New("unsuported file extension~!")
	}
	if fileHeader.Size > maxSize {
		return "", errors.New("file too large (max 5)")
	}
	return "", nil
}
