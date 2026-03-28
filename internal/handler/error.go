package handler

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrInvalidMethod       = errors.New("method not allowed")
	ErrInternalServerError = errors.New("internal server error")
	ErrNilRepo             = errors.New("repository is nil")
)

func MethodNotAllowedError(w http.ResponseWriter) {
	WriteError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, reason error) {
	WriteError(w, reason, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Printf("internal server error: %v", err)

	WriteError(w, ErrInternalServerError, http.StatusInternalServerError)
}
