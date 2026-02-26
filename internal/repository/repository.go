package repository

import (
	"context"

	"movieexample.com/pkg/model"
)

type Repository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	PutUser(ctx context.Context, id string, u *model.User) error
	DeleteUserByID(ctx context.Context, id string) error
}
