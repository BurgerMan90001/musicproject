package memory

import (
	"context"
	"sync"

	"movieexample.com/pkg/model"
)

type Repository struct {
	sync.RWMutex
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	
	return nil, nil
}
func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}
func (r *Repository) PutUser(ctx context.Context, id string, m *model.User) error {
	return nil
}
