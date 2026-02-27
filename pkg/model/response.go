package model

type ErrorResponse struct {
	Message    string `json:"message"`
	Error      error  `json:"error"`
	StatusCode int    `json:"statusCode"`
}
