package models

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type TokenResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}
