package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"github.com/ivegotanidea/golang-gorm-postgres/utils"
	"net/http"
)

func verifyRecaptcha(resp *http.Response, recaptchaVersion string) (bool, float64, error) {
	if recaptchaVersion == "v3" {

		var recaptchaResponse RecaptchaResponseV3
		err := json.NewDecoder(resp.Body).Decode(&recaptchaResponse)
		if err != nil {
			return false, 0, fmt.Errorf("JSON Decode Error: %v", err)
		}

		return recaptchaResponse.Success, recaptchaResponse.Score, nil

	} else {
		var recaptchaResponse RecaptchaResponse
		err := json.NewDecoder(resp.Body).Decode(&recaptchaResponse)
		if err != nil {
			return false, -1, fmt.Errorf("JSON Decode Error: %v", err)
		}

		return recaptchaResponse.Success, 0, nil
	}
}

func RecaptchaMiddleware(secret string, recaptchaVersion string, scoreThreshold float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("g-recaptcha-response")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "reCAPTCHA token missing"})
			return
		}

		resp, err := utils.PostRecaptcha(token, secret)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "reCAPTCHA verification error"})
			return
		}

		success, score, err := verifyRecaptcha(resp, recaptchaVersion)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "reCAPTCHA verification error"})
			return
		}

		if recaptchaVersion == "v3" && (!success || score < scoreThreshold) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "reCAPTCHA verification failed"})
			return
		}

		if recaptchaVersion == "v2" && !success {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "reCAPTCHA verification failed"})
			return
		}

		c.Next()
	}
}
