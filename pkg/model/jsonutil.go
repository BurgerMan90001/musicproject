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

// func Marshal(code int, args ...string) ([]byte, error) {
// 	var message string
// 	if len(args) > 0 {
// 		message = args[0]
// 	}
// 	res := Error{
// 		Code:    code,
// 		Message: message, // fix later
// 	}
// 	return json.Marshal(res)
// }
