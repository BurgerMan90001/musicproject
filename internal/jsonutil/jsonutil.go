package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"musicproject.com/pkg/model"
)

type ContentType string

const Json ContentType = "application/json"

func WriteJSON(w http.ResponseWriter, data any, code int, details ...string) {
	w.Header().Set("Content-type", "application/json")

	jsonEncoder := json.NewEncoder(w)

	switch {
	// Not an error code
	case code >= 200 && code < 300:
		// Check if data is properly encoded
		_, err := json.Marshal(data)
		if err != nil || data == nil {
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

	// Invalid status codes
	case code < 100 || code >= 600:
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Details: append(details, fmt.Sprintf("WriteJSON invalid code: %d", code)),
		})

	default:
		// If the message is an error but the code is error
		err, ok := data.(error)
		if !ok || err == nil || err.Error() == "" {
			w.WriteHeader(http.StatusInternalServerError)
			jsonEncoder.Encode(model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
			})
			return
		}

		if len(details) > 10 {
			details = details[:10]
		}
		w.WriteHeader(code)
		jsonEncoder.Encode(model.ErrorResponse{
			Code:    code,
			Message: err.Error(),
			Details: details,
		})
	}
}
func WriteError(w http.ResponseWriter, reason error, code int, errors ...error) {
	//	if code == http.StatusInternalServerError {
	for _, e := range errors {
		slog.Error(e.Error())
	}
	//}
	WriteJSON(w, reason, code)
}
func ReadJson[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, fmt.Errorf("ReadJSON: %w", err)
	}

	return v, nil
}

// Test helper for reading test responses
func ReadJSONT[T any](t *testing.T, r io.Reader) T {
	t.Helper()
	data, err := ReadJson[T](r)
	require.NoError(t, err)

	return data
}

func MethodNotAllowedError(w http.ResponseWriter) {
	WriteJSON(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, reason error) {
	WriteJSON(w, errors.New("Resource not found"), http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, err error) {
	// TODO USE better logging
	slog.Error("internal server error: ", err)
	WriteError(w, errors.New("Internal server error"), http.StatusInternalServerError)
}
