package handler

import (
	"errors"
	"log"
	"net/http"

	"musicproject.com/internal/jsonutil"
)

var (
	ErrInvalidMethod       = errors.New("method not allowed")
	ErrInternalServerError = errors.New("internal server error")
	//ErrNilRepo             = errors.New("repository is nil")
	ErrInvalidRequestBody  = errors.New("Invalid request body")
)

func MethodNotAllowedError(w http.ResponseWriter) {
	jsonutil.WriteError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, reason error) {
	jsonutil.WriteError(w, reason, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Printf("internal server error: %v", err)
	jsonutil.WriteError(w, err, http.StatusInternalServerError)
}
