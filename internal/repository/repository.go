package repository

import (
	"context"

	"movieexample.com/pkg/model"
)

type Repository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	PutUser(ctx context.Context, id string, u *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

