package repository

import (
	"context"

	"okapi.com/pkg/model"
)

type Repository interface {
	// User methods
	GetUserByID(ctx context.Context, id model.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	PutUser(ctx context.Context, id model.UUID, u *model.User) error
	DeleteUserByID(ctx context.Context, id model.UUID) error

	// Song methods
	GetSongByID(ctx context.Context, id model.UUID) (*model.Song, error)
	PutSong(ctx context.Context, id model.UUID, u *model.Song) error
}
