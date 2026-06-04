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

// ValidateAndUploadFile(file, "./uploads")
func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	ext := strings.ToLower(filepath.Ext((fileHeader.Filename)))

	if !allowExts[ext] {
		return "", errors.New("unsuported file extension~!")
	}
	if fileHeader.Size > maxSize {
		return "", errors.New("file too large (max 5)")
	}

	file, err := fileHeader.Open()
	//
	if err != nil {
		return "", errors.New("Canot open file!")
	}

	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)

	if err != nil {
		return "", errors.New("canot read file")
	}

	mimeType := http.DetectContentType(buffer)

	if !allowMimetypes[mimeType] {
		return "", fmt.Errorf("invalide mimetype:%s", mimeType)
	}
	//Change file name

	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	//Create folder if not exist

	if err = os.MkdirAll(uploadDir, os.ModePerm.Perm()); err != nil {
		return "", errors.New("Canot create upload file")
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
