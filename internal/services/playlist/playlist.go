package playlist

import "github.com/google/uuid"

type Service struct{}

func New() *Service {
	return &Service{}
}

// TODO
func (s *Service) CreatePlaylist(userId uuid.UUID) error {
	return nil
}

func (s *Service) DeletePlaylist(playlistId uuid.UUID) error {
	return nil
}
