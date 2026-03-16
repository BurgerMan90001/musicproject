package handler

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")

	ErrInvalidMethod       = errors.New("method not allowed")
	ErrInternalServerError = errors.New("internal server error")
)
