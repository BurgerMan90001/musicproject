package memory

import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

// Song rating methods
func (r *Repository) GetRatings(ctx context.Context, songId uuid.UUID) ([]model.Rating, error) {
	return nil, nil
}
func (r *Repository) PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error {
	return nil
}
