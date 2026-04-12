package rating

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

type Service struct {
	min  int
	max  int
	repo repository.Rating
}

func New(min, max int, repo repository.Rating) *Service {
	if max == 0 {
		max = 5
	}
	return &Service{
		repo: repo,
		min:  min,
		max:  max,
	}
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
		sum += rating.Value
	}
	return sum / float64(len(ratings)), nil
}

func (s *Service) Put(ctx context.Context, rating *model.Rating) error {
	//rating, err := s.repo.GetRating(ctx, rating.SongID, rating.UserID)
	// A rating exists

	if rating.Value < float64(s.min) {
		return fmt.Errorf("Invalid rating: %v rating too small", rating.Value)
	}
	if rating.Value > float64(s.max) {
		return fmt.Errorf("Invalid rating: %v rating too big", rating.Value)
	}
	// Update rating if it exists
	if err := s.repo.Update(ctx, rating); !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	_, err := s.repo.Put(ctx, rating)
	return err
}
