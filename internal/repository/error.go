package repository

import (
	"net/http"

	"songsled.com/pkg/model"
)

var (
	ErrNotFound = &model.Error{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
	ErrUserNotFound = &model.Error{
		Code:    http.StatusNotFound,
		Message: "User not found",
	}
	ErrEmailTaken = &model.Error{
		Code:    http.StatusConflict,
		Message: "Email is already taken",
	}
	ErrSongNotFound = &model.Error{
		Code:    http.StatusNotFound,
		Message: "Song not found",
	}
)
