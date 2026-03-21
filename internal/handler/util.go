package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"musicproject.com/pkg/model"
)

type HandlerTest struct {
	Name   string
	Method string
	Body   map[string]any

	WantCode int

	WantData    any
	WantMessage string

	RepoItem any
	RepoErr  error
}

func WriteJSON(w http.ResponseWriter, data any, code int, args ...string) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	var en []byte
	var err error

	if data != nil {
		en, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	var message string
	if len(args) > 0 {
		message = args[0]
	}

	res := model.Response{
		Success: code >= 200 && code < 300,
		Data:    en,
		Message: message, // fix later
	}
	return json.NewEncoder(w).Encode(res)
}

func WriteError(w http.ResponseWriter, reason error, code int) error {
	err := WriteJSON(w, nil, code, reason.Error())
	return err
}

func NewRequestBody(body any) (io.Reader, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(buf), nil
}

func MethodNotAllowedError(w http.ResponseWriter, r *http.Request) {
	WriteError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, r *http.Request, reason error) {
	WriteError(w, reason, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %v", err)

	WriteError(w, ErrInternalServerError, http.StatusInternalServerError)
}
