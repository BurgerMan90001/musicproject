package song

import (
	"musicproject.com/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

/*
func (c *Service) GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error) {
	song, err := c.repo.GetSongByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return song, nil
}
func (c *Service) GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error) {
	songs, err := c.repo.GetSongsByGenre(ctx, genre)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return songs, nil
}
*/
