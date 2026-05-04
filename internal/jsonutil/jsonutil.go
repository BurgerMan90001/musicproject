package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"songsled.com/pkg/model"
)

// type ContentType string

// const Json ContentType = "application/json"

// Sets header status code
func WriteJSON(w http.ResponseWriter, data any, code int, details ...string) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)

	switch {
	case data == nil:
		WriteError(w, &model.Error{
			Code:    code,
			Message: "Internal server error",
			Details: fmt.Sprintf("%v; WriteJSON: data is nil", details),
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
				Details: fmt.Sprintf("%v; WriteJSON: marshal data error %v", details, err),
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
			Details: fmt.Sprintf("%v; WriteJSON invalid code: %d", details, code),
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
	jerr := newJsonErr(reason)

	if os.Getenv("SHOW_ERROR_DETAILS") != "true" {
		jerr.Details = ""
	}
	// Service error / internal error
	if jerr.Code >= 500 && jerr.Code < 600 {
		w.WriteHeader(jerr.Code)
		

		json.NewEncoder(w).Encode(jerr)

		slog.Error(fmt.Sprintf("%v", reason))
		return
	}

	w.WriteHeader(jerr.Code)
	json.NewEncoder(w).Encode(jerr)
}

// TODO have better details in json response when an internal error occurs
func newJsonErr(reason error) *model.Error {

	var jerr *model.Error
	if errors.As(reason, &jerr) {
		switch {
		case jerr == nil:
			jerr = &model.Error{}
			// if reason != nil {
			// 	jerr.Details = reason.Error()
			// }
		case jerr.Message == "":
			jerr.Details = "WriteError: empty error message"

		case jerr.Code < 100 || jerr.Code >= 600:
			jerr.Details = fmt.Sprintf("WriteError: invalid status code %d", jerr.Code)
		// Error is not an internal server error
		default:
			return jerr
		}
	} else {
		jerr = &model.Error{}
		if reason != nil {
			jerr.Details = reason.Error()
		}
	}
	jerr.Code = http.StatusInternalServerError
	jerr.Message = "Internal server error"

	return jerr
}

func ReadJson[T any](r io.Reader) (T, error) {
	var v T

	data, err := io.ReadAll(r)
	if err != nil {
		return v, err
	}
	if len(data) == 0 {
		return v, nil
	}

	if err := json.Unmarshal(data, &v); err != nil {

		jerr := &model.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Body is not of type %T %s", *new(T), string(data)),
			Details: fmt.Sprintf("ReadJson: %T\n Reason: %s\n JSON: %s", *new(T), err.Error(), string(data)),
		}
		// if testing.Testing() {
		// 	slog.Info(jerr.Details)
		// }

		return v, jerr
	}

	return v, nil
}

func WriteNotFound(w http.ResponseWriter, reason error) {
	err := &model.Error{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
	if reason != nil {
		err.Details = reason.Error()
	}
	WriteError(w, err)
}

func WriteInvalidRequestBody(w http.ResponseWriter, reason error) {
	err := &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid request body",
	}
	if reason != nil {
		err.Details = reason.Error()
	}
	WriteError(w, err)
}

// Responds with generic internal server error json response
// Logs error
func InternalServerError(w http.ResponseWriter, reason error) {

	slog.Error("Internal server error: ", reason.Error(), "")
	WriteError(w, &model.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	})
}

func NotImplemented(w http.ResponseWriter) {
	WriteError(w, &model.Error{
		Code:    http.StatusNotImplemented,
		Message: "Route is not implemented yet!",
	})
}
