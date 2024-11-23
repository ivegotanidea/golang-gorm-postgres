package utils

import (
	"net/http"
	"net/url"
)

func PostRecaptcha(token string, secret string) (*http.Response, error) {
	recaptchaURL := "https://www.google.com/recaptcha/api/siteverify"

	data := url.Values{}
	data.Set("secret", secret)
	data.Set("response", token)

	return http.PostForm(recaptchaURL, data)
}
