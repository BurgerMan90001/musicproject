package rating

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAggregatedRating(ctx context.Context, songId uuid.UUID) (float64, error) {
	ratings, err := s.repo.GetRatings(ctx, songId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, repository.ErrNotFound
		}
		return 0, err
	}
	var sum float64 = 0
	for _, rating := range ratings {
		sum += float64(rating.Value)
	}
	aggregatedRating := sum / float64(len(ratings))
	return aggregatedRating, nil
}

/*
func (c *Service) PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error {
	return c.repo.PutRating(ctx, songId, userId, value)
}
*/
