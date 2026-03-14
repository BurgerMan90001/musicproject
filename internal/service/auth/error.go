package auth

import "errors"

var (
	ErrInvalidPassword = errors.New("password must be at least 8 characters, contain uppercase, lowercase, number, and special character")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidToken    = errors.New("invalid token")
)
