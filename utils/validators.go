package utils

import (
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateImageURL(fl validator.FieldLevel) bool {
	// Get the value of the field
	imageURL := fl.Field().String()

	// Parse the URL
	parsedURL, err := url.ParseRequestURI(imageURL)
	if err != nil {
		return false
	}

	// Check if the URL ends with a valid image file extension
	validExtensions := []string{".jpg", ".jpeg"}
	for _, ext := range validExtensions {
		if strings.HasSuffix(strings.ToLower(parsedURL.Path), ext) {
			return true
		}
	}

	return false
}
