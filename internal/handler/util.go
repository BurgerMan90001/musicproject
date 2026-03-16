package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"musicproject.com/pkg/model"
)

const (
	StatusSucess = "success"
	StatusError  = "error"
)

func NewRequestBody(body any) (io.Reader, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(buf), nil
}

func ReadJSON[T any](r io.ReadCloser) (T, error) {
	var v T

	err := json.NewDecoder(r).Decode(&v)
	return v, errors.Join(err, r.Close())
}

func WriteJSON(w http.ResponseWriter, status string, data any, code int, args ...string) error {
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
		Status:  status,
		Data:    en,
		Message: message, // fix later
	}
	return json.NewEncoder(w).Encode(res)
}
func MarshalJSON(status string, data any, code int, args ...string) ([]byte, error) {
	var en []byte
	var err error

	if data != nil {
		en, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	var message string
	if len(args) > 0 {
		message = args[0]
	}
	res := model.Response{
		Status:  status,
		Data:    en,
		Message: message, // fix later
	}
	return json.Marshal(res)
}
func WriteErrJSON(w http.ResponseWriter, reason error, code int) error {
	err := WriteJSON(w, StatusError, nil, code, reason.Error())

	return err
}

func MethodNotAllowedError(w http.ResponseWriter, r *http.Request) {
	WriteErrJSON(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, r *http.Request, reason error) {
	WriteErrJSON(w, reason, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %v", err)

	WriteErrJSON(w, ErrInternalServerError, http.StatusInternalServerError)
}

type HandlerTest struct {
	Name   string
	Method string
	Body   map[string]any

	WantCode int

	WantStatus  string
	WantData    any
	WantMessage string

	RepoItem any
	RepoErr  error
}
