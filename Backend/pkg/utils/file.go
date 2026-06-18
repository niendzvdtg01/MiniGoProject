package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var allowMimetypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

const maxSize = 5 << 20

func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	if !allowExts[ext] {
		return "", errors.New("unsupported file extension")
	}
	if fileHeader.Size > maxSize {
		return "", errors.New("file too large (max 5 MB)")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.New("cannot open file")
	}

	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		return "", errors.New("cannot read file")
	}

	mimeType := http.DetectContentType(buffer)

	if !allowMimetypes[mimeType] {
		return "", fmt.Errorf("invalid mimetype: %s", mimeType)
	}

	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	if err = os.MkdirAll(uploadDir, 0o755); err != nil {
		return "", errors.New("cannot create upload directory")
	}

	savePath := filepath.Join(uploadDir, filename)

	if err := saveFile(fileHeader, savePath); err != nil {
		return "", err
	}

	return filename, nil
}

func saveFile(fileHeader *multipart.FileHeader, destination string) error {
	src, err := fileHeader.Open()

	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(destination)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, src)

	if err != nil {
		return err
	}

	return nil
}
