package handler

import (
	"encoding/json"
	"net/http"

	"musicproject.com/pkg/model"
)

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
