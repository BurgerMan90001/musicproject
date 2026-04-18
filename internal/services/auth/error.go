package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"musicproject.com/internal/services"
)

var (
	ErrInvalidPassword  = services.NewErr("Password must be at least 8 characters, contain uppercase, lowercase, number, and special character")
	ErrMismatchPassword = services.NewErr("Incorrect password")
	ErrInvalidEmail     = services.NewErr("Invalid email")
	ErrIncorrectLogin   = services.NewErr("Incorrect password or email")

	ErrNoAccessToken    = services.NewErr("No access token cookie")
	ErrNoRefeshToken    = services.NewErr("No refresh token cookie")
	ErrTokenExpired     = jwt.ErrTokenExpired
	ErrInvalidTokenType = services.NewErr("Invalid token type")

	ErrUserAlreadyExists = services.NewErr("User already exists")
)
