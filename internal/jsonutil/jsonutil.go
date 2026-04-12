package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"musicproject.com/pkg/model"
)

func WriteJSON(w http.ResponseWriter, data any, code int, details ...string) {
	w.Header().Set("Content-type", "application/json")

	jsonEncoder := json.NewEncoder(w)

	// Check for invalid status code
	if http.StatusText(code) == "" {
		// Is internal server error
		invaldCode := code
		code = http.StatusInternalServerError

		w.WriteHeader(code)
		jsonEncoder.Encode(model.ErrorResponse{
			Code:    code,
			Message: "Internal server error",
			Details: append(details, fmt.Sprintf("WriteJSON invalid code: %d", invaldCode)),
		})
		return
	}

	var message string = ""
	if len(details) > 0 {
		message = details[0]
		details = details[1:]
		if len(details) > 10 {
			details = details[1:10]
		}
	}

	success := code >= 200 && code < 300

	if success {
		// Check if data is properly encoded
		_, err := json.Marshal(data)
		if err != nil {
			code = http.StatusInternalServerError

			w.WriteHeader(code)
			jsonEncoder.Encode(model.ErrorResponse{
				Code:    code,
				Message: "Internal server error",
				Details: details,
			})
			return
		}
		w.WriteHeader(code)
		jsonEncoder.Encode(data)
		return
	}

	w.WriteHeader(code)
	jsonEncoder.Encode(model.ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	})
}

func WriteError(w http.ResponseWriter, reason error, code int) {
	WriteJSON(w, nil, code, reason.Error())
}

func ReadJSON[T any](r io.ReadCloser) (T, error) {
	var v T
	err := json.NewDecoder(r).Decode(&v)

	return v, errors.Join(err, r.Close())
}

// Test helper for reading test responses
func ReadJSONT[T any](t *testing.T, r io.ReadCloser) T {
	t.Helper()
	data, err := ReadJSON[T](r)
	require.NoError(t, err)

	return data
}

// TODO update to match WriteJSON
func Marshal(code int, args ...string) ([]byte, error) {
	// var message string
	// if len(args) > 0 {
	// 	message = args[0]
	// }
	// // res := Error{
	// // 	Code:    code,
	// // 	Message: message, // fix later
	// // }
	// return json.Marshal(res)
	return nil, nil
}

func MethodNotAllowedError(w http.ResponseWriter) {
	WriteError(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, reason error) {
	WriteError(w, errors.New("Resource not found"), http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, err error) {
	// TODO USE better logging
	log.Printf("internal server error: %v", err)
	WriteError(w, err, http.StatusInternalServerError)
}
