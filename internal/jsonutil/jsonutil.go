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

// type ContentType string

// const Json ContentType = "application/json"

// Sets header status code
func WriteJSON(w http.ResponseWriter, data any, code int, details ...string) {
	w.Header().Set("Content-type", "application/json")
	jsonEncoder := json.NewEncoder(w)

	// if len(details) > 10 {
	// 	details = details[:10]
	// }

	switch {
	case data == nil:

		WriteError(w, &model.Error{
			Code:    code,
			Message: "Internal server error",
			Details: append(details, "WriteJSON: data is nil"),
		})
	// Not an error code
	case code >= 200 && code < 300:
		// Check if data is properly encoded
		_, err := json.Marshal(data)
		if err != nil {
			code = http.StatusInternalServerError

			WriteError(w, &model.Error{
				Code:    code,
				Message: "Internal server error",
				Details: append(details, fmt.Sprintf("WriteJSON: marshal data error %v", err)),
			})
			return
		}
		w.WriteHeader(code)
		jsonEncoder.Encode(data)

	// Invalid status codes
	case code < 100 || code >= 600:
		WriteError(w, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Details: append(details, fmt.Sprintf("WriteJSON invalid code: %d", code)),
		})

	default:
		err := data.(error)
		WriteError(w, err)
	}
}
func validate(merr *model.Error, ok bool) error {

	var m string
	switch {
	case !ok:
		m = "WriteError: reason is not of type model.Error"
	case merr == nil || merr.Message == "":
		m = "WriteError: empty error message"
	case merr.Code < 100 || merr.Code >= 600:
		m = "WriteError: reason is not of type model.Error"
	}
	if m == "" {
		return merr
	}
	return errors.New(m)
}

// Writes json response error
// Use model.Error for status codes and error details
// Sets header status code
func WriteError(w http.ResponseWriter, reason error) {
	w.Header().Set("Content-type", "application/json")
	merr, ok := reason.(*model.Error)
	err := validate(merr, ok)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Error{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Details: []string{err.Error()},
		})
		slog.Error(err.Error())
		return
	}

	// Check if service error
	if merr.Code >= 500 && merr.Code < 600 {
		for _, e := range merr.Details {
			slog.Error(e)
		}
	}
	w.WriteHeader(merr.Code)
	json.NewEncoder(w).Encode(merr)

}
func ReadJson[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Details: []string{err.Error()},
		}
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
	WriteError(w, &model.Error{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	})
}

func NotFoundError(w http.ResponseWriter, reason error) {
	WriteJSON(w, errors.New("Resource not found"), http.StatusNotFound)
}

// Responds with generic internal server error json response
// Logs error
func InternalServerError(w http.ResponseWriter, err error) {
	// TODO USE better logging
	slog.Error("internal server error: ", err.Error(), "")
	WriteError(w, &model.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	})
}
