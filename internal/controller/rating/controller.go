package rating

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
)

type Controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetAggregatedRating(ctx context.Context, songId uuid.UUID) (float64, error) {
	ratings, err := c.repo.GetRatings(ctx, songId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, repository.ErrNotFound
		}
		return 0, err
	}
	var sum float64 = 0
	for _, rating := range ratings {
		sum += rating.Value
	}
	return sum, nil
}

func (c *Controller) PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error {
	
	return nil
}
