package search

import (
	"context"

	"musicproject.com/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Filter(ctx context.Context, query ...string) {

}
