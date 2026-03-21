package model

import (
	"encoding/json"
	"errors"
	"io"
)

func ReadJSON[T any](r io.ReadCloser) (T, error) {
	var v T

	err := json.NewDecoder(r).Decode(&v)
	return v, errors.Join(err, r.Close())
}

func MarshalJSON(data any, code int, args ...string) ([]byte, error) {
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
	res := Response{
		Success: code >= 200 && code < 300,
		Data:    en,
		Message: message, // fix later
	}
	return json.Marshal(res)
}
