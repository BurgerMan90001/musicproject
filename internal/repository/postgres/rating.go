package postgres

import (
	"context"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
)

type Rating struct {
	q *gensqlc.Queries
}

func NewRating(q *gensqlc.Queries) *Rating {
	return &Rating{q}
}

// TODO FIX rating conversions
// Song rating methods
func (r *Rating) GetAggregatedRating(ctx context.Context, songId uuid.UUID) (float64, error) {
	ratings, err := r.q.GetAggregatedRating(ctx, songId)
	if err != nil {
		return 0, err
	}

	return float64(ratings), nil
}
func (r *Rating) PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error {
	return r.q.PutRating(ctx, gensqlc.PutRatingParams{
		SongID: songId, RatingValue: int32(value),
	})
}
