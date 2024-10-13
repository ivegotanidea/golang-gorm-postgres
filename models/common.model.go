package models

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type SuccessPageResponse struct {
	Status  string `json:"status"`
	Data    any    `json:"data"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

type TokenResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}
