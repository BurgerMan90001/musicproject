package memory

import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

func (r *Repository) GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error) {
	return nil, nil
}
func (r *Repository) GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error) {
	return nil, nil
}
func (r *Repository) PutSong(ctx context.Context, id uuid.UUID, u *model.Song) error {
	return nil
}
