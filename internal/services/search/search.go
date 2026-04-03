package search

import "context"

type Service interface {
	Filter(ctx context.Context) error
}

// func New(repo repository.Song) *Service {

// 	return &Service{repo: repo}
// }

// func (s *Service) Filter(ctx context.Context, query ...string) {

// }
