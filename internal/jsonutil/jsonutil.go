package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

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

// Writes json response error
// Use model.Error for status codes and error details
// Sets header status code
func WriteError(w http.ResponseWriter, reason error) {
	w.Header().Set("Content-type", "application/json")
	merr := validateErr(reason)
	// Service error / internal error
	if merr.Code >= 500 && merr.Code < 600 {
		w.WriteHeader(merr.Code)
		// When not testing
		// if !testing.Testing() {
		// 	merr.Details = nil
		// }
		json.NewEncoder(w).Encode(merr)

		slog.Error(fmt.Sprintf("%v", reason))
		return
	}

	w.WriteHeader(merr.Code)
	json.NewEncoder(w).Encode(merr)
}

// TODO have better details in json response when an internal error occurs
func validateErr(reason error) *model.Error {
	merr, ok := reason.(*model.Error)

	switch {
	case merr == nil || !ok:
		// Create new error
		merr = &model.Error{}
		merr.Details = append(merr.Details, "WriteError: reason is not of type model.Error")

	case merr.Message == "":
		merr.Details = append(merr.Details, "WriteError: empty error message")

	case merr.Code < 100 || merr.Code >= 600:
		merr.Details = append(merr.Details, "WriteError: invalid status code")
	default:
		return merr
	}
	merr.Code = http.StatusInternalServerError
	merr.Message = "Internal server error"
	return merr
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

func WriteMethodNotAllowed(w http.ResponseWriter) {
	WriteError(w, &model.Error{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	})
}

func WriteNotFound(w http.ResponseWriter, reason error) {
	err := &model.Error{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
		//Details: []string{reason.Error()},
	}
	if reason != nil {
		err.Details = append(err.Details, reason.Error())
	}
	WriteError(w, err)
}

func WriteInvalidRequestBody(w http.ResponseWriter, reason error) {
	WriteError(w, &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid request body",
		Details: []string{reason.Error()},
	})
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
