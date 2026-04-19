package repository

import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

// TODO NOT USED
//go:generate mockgen -destination=../../gen/mocks/repository.go -package=mocks  -source=repository.go

type Repo[T any] interface {
	GetByID(ctx context.Context, id uuid.UUID) (T, error)
	Put(ctx context.Context, item T) (uuid.UUID, error)
	DeleteByID(ctx context.Context, userId uuid.UUID) error
}
type User interface {
	Repo[*model.User]
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}
type Rating interface {
	Put(ctx context.Context, rating *model.Rating) (uuid.UUID, error)
	GetRatings(ctx context.Context, songId uuid.UUID) ([]*model.Rating, error)
	//GetRating(ctx context.Context, songId, userId uuid.UUID) (*model.Rating, error)
	Update(ctx context.Context, rating *model.Rating) error
}

// Metadata repository for songs
type Song interface {
	Repo[*model.Song]
	GetSongsByGenre(ctx context.Context, genre string) ([]*model.Song, error)
}

type Token interface {
	RevokeToken(ctx context.Context, tokenId uuid.UUID) error
	Revoked(ctx context.Context, tokenId uuid.UUID) error
}
