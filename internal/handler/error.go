package handler

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")

	ErrInvalidMethod       = errors.New("method not allowed")
	ErrInternalServerError = errors.New("internal server error")
)
