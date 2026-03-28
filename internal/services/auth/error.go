package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidPassword  = errors.New("password must be at least 8 characters, contain uppercase, lowercase, number, and special character")
	ErrMismatchPassword = errors.New("incorrect password")
	ErrInvalidEmail     = errors.New("invalid email")

	ErrInvalidToken            = errors.New("invalid token")
	ErrNoAccessToken           = errors.New("no access token cookie")
	ErrNoRefeshToken           = errors.New("no access token cookie")
	ErrTokenExpired            = jwt.ErrTokenExpired
	ErrInvalidTokenType        = errors.New("invalid token type")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

	ErrUserAlreadyExists = errors.New("user already exists")
)
