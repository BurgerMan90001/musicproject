package repository

import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

type Repository interface {
	// User methods
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	PutUser(ctx context.Context, id uuid.UUID, u *model.User) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error

	// Song methods
	GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error)
	GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error)
	PutSong(ctx context.Context, id uuid.UUID, u *model.Song) error

	// Song racting methods
	GetRatings(ctx context.Context, songId uuid.UUID) ([]model.Rating, error)
	PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error

	Stop() error
}
