package auth

import (
	"net/http"

	"musicproject.com/pkg/model"
)

var (
	ErrInvalidPassword = &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Password must be at least 8 characters, contain uppercase, lowercase, number, and special character",
	}
	ErrMismatchPassword = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Incorrect password",
	}
	ErrInvalidEmail = &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid email",
	}
	ErrIncorrectLogin = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Incorrect password or email",
	}

	ErrNoAccessToken = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "No access token",
	}
	ErrInvalidClaims = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Invalid claims",
	}
	ErrTokenExpired = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Token is expired",
	}

	ErrInvalidTokenType = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Invalid token type",
	}
	ErrUserAlreadyExists = &model.Error{
		Code:    http.StatusConflict,
		Message: "User already exists",
	}
)
