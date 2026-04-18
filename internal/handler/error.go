package handler

import (
	"errors"
)

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrInvalidRequestBody  = errors.New("Invalid request body")
)
