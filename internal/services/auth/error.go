package auth

import (
	"net/http"

	"songsled.com/pkg/model"
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
	ErrIncorrectLogin = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Incorrect password or email",
	}
)

// Reason is optional
func ErrInvalidToken(reasons ...error) *model.Error {
	jerr := &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Invalid token",
	}
	for _, r := range reasons {
		if r != nil {
			jerr.Details += r.Error()
		}
	}

	return jerr
}

// ErrNoToken = &model.Error{
// 	Code:    http.StatusUnauthorized,
// 	Message: "No token present",
// }
// ErrTokenRevoked = &model.Error{
// 	Code:    http.StatusUnauthorized,
// 	Message: "Token revoked",
// }
// ErrInvalidClaims = &model.Error{
// 	Code:    http.StatusUnauthorized,
// 	Message: "Invalid claims",
// }
// ErrTokenExpired = &model.Error{
// 	Code:    http.StatusUnauthorized,
// 	Message: "Token is expired",
// }

// ErrInvalidTokenType = &model.Error{
// 	Code:    http.StatusUnauthorized,
// 	Message: "Invalid token type",
// }
