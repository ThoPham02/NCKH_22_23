package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ErrResponse(err error) gin.H {
	return gin.H{"error": err}
}

func ErrRequest(err error) gin.H {
	var errMessage string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMessage = fmt.Sprintf("%s is required", err.Field())
		case "email":
			errMessage = fmt.Sprintf("%s is not a valid email address", err.Field())
		case "min":
			errMessage = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
		case "max":
			errMessage = fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
		case "gte":
			errMessage = fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
		case "lte":
			errMessage = fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
		}
	}
	return gin.H{"error": errMessage}
}

func ErrResourceNotFound() gin.H {
	return gin.H{"error": "Resource not found"}
}
