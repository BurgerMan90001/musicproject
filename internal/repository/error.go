package repository

import (
	"net/http"

	"musicproject.com/pkg/model"
)

var (
	ErrNotFound = &model.Error{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
	ErrTokenRevoked = &model.Error{
		Code:    http.StatusUnauthorized,
		Message: "Token revoked",
	}
)
