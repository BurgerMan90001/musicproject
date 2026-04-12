package handler

import (
	"errors"
)

var (
	ErrInvalidMethod       = errors.New("Method not allowed")
	ErrInternalServerError = errors.New("Internal server error")
	ErrInvalidRequestBody  = errors.New("Invalid request body")

	ErrUnauthorized = errors.New("Unauthorized")
)
