package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			log.Printf("%s", e.Namespace())
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")
			for i, part := range parts {
				if strings.Contains(part, "[") {
					idx := strings.Index(part, "[")
					base := camelToSnake(part[:idx])
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
				errors[fieldPath] = fmt.Sprintf("%s must be smaller than: %s", fieldPath, e.Param())
			case "min_int":
				errors[fieldPath] = fmt.Sprintf("%s must be bigger than: %s", fieldPath, e.Param())
			case "file_extension":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[fieldPath] = fmt.Sprintf("%s only accept the file have extension %s", fieldPath, allowedValues)
			case "email_advance":
				errors[fieldPath] = fmt.Sprintf("email stay in banned list:%s", fieldPath)
			case "password_strong":
				errors[fieldPath] = fmt.Sprintf("password must have valide format:%s", fieldPath)
			}

		}
		return gin.H{"error": errors}
	}
	return gin.H{"error": "invalid request: " + err.Error()}
}

func RegisterValidation(v *validator.Validate) error {

	var blockDomains = map[string]bool{
		"blacklist.com": true,
		"edu.vn":        true,
		"abc.com":       true,
	}

	v.RegisterValidation("email_advance", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()

		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			return false
		}

		domain := NormalizeString(parts[1])
		return !blockDomains[domain]

	})

	//password validation
	v.RegisterValidation("password_strong", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 {
			return false
		}

		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpeChar := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>_\-+=\\[\]/~]`).MatchString(password)

		return hasLower && hasUpper && hasDigit && hasSpeChar
	})

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
