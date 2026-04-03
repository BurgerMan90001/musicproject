package handler

import (
	"encoding/json"
	"net/http"

	"musicproject.com/pkg/model"
)

func WriteJSON(w http.ResponseWriter, data any, code int, args ...string) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	var res any
	success := code >= 200 && code < 300

	if success {
		res = data
	} else {
		var (
			message string
			details []string
		)
		if len(args) > 0 {
			message = args[0]
			if len(args) > 10 {
				details = args[1:10]
			} else {
				details = args[1:]
			}
		}

		res = model.Error{
			Code:    code,
			Message: message,
			Details: details,
		}
	}
	return json.NewEncoder(w).Encode(res)
}

func WriteError(w http.ResponseWriter, reason error, code int) error {
	err := WriteJSON(w, nil, code, reason.Error())
	return err
}
