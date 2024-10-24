package models

// ErrorResponse represents a general error response.
// @Description This model is returned when an API call fails due to an error.
type ErrorResponse struct {
	// Status represents the status of the response, usually set to "error".
	// Example: "error"
	Status string `json:"status"`
	// Message provides a human-readable description of the error that occurred.
	// Example: "Invalid request payload"
	Message string `json:"message"`
}

// SuccessResponse represents a successful response.
// @Description This model is returned when an API call succeeds.
type SuccessResponse struct {
	// Status represents the status of the response, typically set to "success".
	// Example: "success"
	Status string `json:"status"`
	// Data contains the result of the successful operation. Can be any type of data.
	// Example: {"id": 1, "name": "John Doe"}
	Data any `json:"data"`
}

// SuccessPageResponse represents a paginated response.
// @Description This model is used when paginated data is returned from the API.
type SuccessPageResponse struct {
	// Status represents the status of the response, typically set to "success".
	// Example: "success"
	Status string `json:"status"`
	// Data contains the data payload for the current page. Can be any type of data.
	// Example: [{"id": 1, "name": "Item 1"}, {"id": 2, "name": "Item 2"}]
	Data any `json:"data"`
	// Results specifies the number of items returned in the current page.
	// Example: 10
	Results int `json:"results"`
	// Page specifies the current page number in the paginated result set.
	// Example: 1
	Page int `json:"page"`
	// Limit specifies the maximum number of items that can be returned in a single page.
	// Example: 10
	Limit int `json:"limit"`
}

// TokenResponse represents a token response, usually after successful authentication.
// @Description This model is used to return an access token after a user logs in or when a token is refreshed.
type TokenResponse struct {
	// Status represents the status of the response, typically set to "success".
	// Example: "success"
	Status string `json:"status"`
	// AccessToken is the token that allows the user to authenticate subsequent API requests.
	// Example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
	AccessToken string `json:"access_token"`
}
