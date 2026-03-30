package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidPassword  = errors.New("Password must be at least 8 characters, contain uppercase, lowercase, number, and special character")
	ErrMismatchPassword = errors.New("Incorrect password")
	ErrInvalidEmail     = errors.New("Invalid email")
	ErrIncorrectLogin   = errors.New("Incorrect password or email")

	ErrInvalidToken            = errors.New("Invalid token")
	ErrNoAccessToken           = errors.New("No access token cookie")
	ErrNoRefeshToken           = errors.New("No access token cookie")
	ErrTokenExpired            = jwt.ErrTokenExpired
	ErrInvalidTokenType        = errors.New("Invalid token type")
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")

	ErrUserAlreadyExists = errors.New("User already exists")
)
