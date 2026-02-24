package repository

import (
	"context"

	"movieexample.com/pkg/model"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	PutUser(ctx context.Context, id string, u *model.User) error
}
