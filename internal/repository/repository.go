package repository

import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"

	_ "github.com/lib/pq"
)

type User interface {
	GetUserByID(ctx context.Context, userId uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	PutUser(ctx context.Context, email string, passwordHash string) (uuid.UUID, error)
	DeleteUserByID(ctx context.Context, userId uuid.UUID) error
}

type Rating interface {
	GetRatings(ctx context.Context, songId uuid.UUID) ([]model.Rating, error)
	PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) (uuid.UUID, error)
}

type Song interface {
	GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Song, error)
	GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error)
	PutSong(ctx context.Context, songId uuid.UUID, song *model.Song) (uuid.UUID, error)
}
