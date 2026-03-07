package song

import (
	"context"
	"errors"

	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

type Controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error) {
	songs, err := c.repo.GetSongsByGenre(ctx, genre)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return songs, nil
}
