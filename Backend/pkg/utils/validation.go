package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			log.Printf("%s", e.Namespace())
			//split each attribute
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")
			for i, part := range parts {
				if strings.Contains("part", "[") {
					idx := strings.Index(part, "[")
					base := camelToSnake(part[:idx]) // 0 den truoc dau vuong [
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = camelToSnake(part)
				}
			}

			fieldPath := strings.Join(parts, ".")

			switch e.Tag() {
			case "required":
				errors[fieldPath] = fmt.Sprintf("%s is a must", fieldPath)
			case "gt":
				errors[fieldPath] = fieldPath + " the number must be larger than zero!"
			case "uuid":
				errors[fieldPath] = fieldPath + " the number must be a valid uuid!"
			case "slug":
				errors[fieldPath] = fieldPath + " Only have normal letters, numbers...!"
			case "min":
				errors[fieldPath] = fmt.Sprintf(
					"%s must have at least %s characters",
					fieldPath,
					e.Param(),
				)

			case "max":
				errors[fieldPath] = fmt.Sprintf(
					"%s must not exceed %s characters",
					fieldPath,
					e.Param(),
				)
			case "search":
				errors[fieldPath] = fieldPath + " contains invalid characters"
			case "max_int":
				errors[fieldPath] = fmt.Sprintf("%s must be smaller than: %s", fieldPath, e.Tag())
			case "min_int":
				errors[fieldPath] = fmt.Sprintf("%s must be bigger than: %s", fieldPath, e.Tag())
			case "file_extension":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[fieldPath] = fmt.Sprintf("%s only accept the file have extension %s", fieldPath, allowedValues)
			}

		}
		return gin.H{"error": errors}
	}
	return gin.H{"error": "Invalid request: " + err.Error()}
}

func RegisterValidation() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return fmt.Errorf("Fail to get engine validator")
	}

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() >= minVal
	})
	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() <= maxVal
	})
	//file extension:jpg, mp4...
	v.RegisterValidation("file_extension", func(fl validator.FieldLevel) bool {
		fileName := fl.Field().String()
		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}
		allowExtension := strings.Fields(allowedStr)

		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(fileName)), ".")

		for _, allowed := range allowExtension {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}
		return false
	})
	return nil
}
