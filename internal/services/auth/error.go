package auth

import (
	"net/http"

	"musicproject.com/pkg/model"
)

var (
	// Signup errors
	ErrInvalidPassword = &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Password must be at least 8 characters, contain uppercase, lowercase, number, and special character",
	}
	ErrInvalidEmail = &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid email",
	}

	// Login errors
	ErrMismatchPassword = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Incorrect password",
	}

	ErrIncorrectLogin = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Incorrect password or email",
	}

	ErrNoToken = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "No token present",
	}
	ErrTokenRevoked = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Token revoked",
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
