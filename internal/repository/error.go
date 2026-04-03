package repository

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrNilRepo  = errors.New("repo in context is nil")
)
