package repository

//go:generate mockgen -destination=../../gen/mocks/repository.go -package=mocks  -source=repository.go
import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

type Repository interface {
	GetUserByID(ctx context.Context, userId uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	PutUser(ctx context.Context, email string, passwordHash string) (uuid.UUID, error)
	DeleteUserByID(ctx context.Context, userId uuid.UUID) error

	GetRatings(ctx context.Context, songId uuid.UUID) ([]model.Rating, error)
	PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error

	GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Song, error)
	GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error)
	PutSong(ctx context.Context, songId uuid.UUID, song *model.Song) (uuid.UUID, error)

	SearchSongs(ctx context.Context) ([]model.Song, error)

	GetToken(ctx context.Context, tokenString string) error
	PutToken(ctx context.Context, tokenString string) error
	DeleteToken(ctx context.Context, tokenString string) error
}
