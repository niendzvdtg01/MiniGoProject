package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " the number must be larger than zero!"
			case "uuid":
				errors[e.Field()] = e.Field() + " the number must be a valid uuid!"
			case "slug":
				errors[e.Field()] = e.Field() + " Only have normal letters, numbers...!"
			case "min":
				errors[e.Field()] = fmt.Sprintf(
					"%s must have at least %s characters",
					e.Field(),
					e.Param(),
				)

			case "max":
				errors[e.Field()] = fmt.Sprintf(
					"%s must not exceed %s characters",
					e.Field(),
					e.Param(),
				)
			case "search":
				errors[e.Field()] = e.Field() + " contains invalid characters"
			case "max_int":
				errors[e.Field()] = fmt.Sprintf("%s must be smaller than: %s", e.Field(), e.Tag())
			case "min_int":
				errors[e.Field()] = fmt.Sprintf("%s must be bigger than: %s", e.Field(), e.Tag())
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
	return nil
}
