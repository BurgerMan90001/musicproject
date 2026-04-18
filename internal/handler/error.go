package handler

import (
	"errors"
)

var _ error = (*HandlerErr)(nil)
var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrInvalidRequestBody  = errors.New("Invalid request body")
)

type HandlerErr struct {
	s string
}

func (e *HandlerErr) Error() string {
	return e.s
}
