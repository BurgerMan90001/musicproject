package repository

import (
	"net/http"

	"musicproject.com/pkg/model"
)

var (
	ErrNotFound = &model.Error{
		Code:    http.StatusNotFound,
		Message: "not found",
	}
)
