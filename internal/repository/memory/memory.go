package memory

import (
	"context"
	"sync"

	"okapi.com/internal/repository"
	"okapi.com/pkg/model"
)

type Repository struct {
	sync.RWMutex
	users map[string]*model.User
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	r.RLock()
	defer r.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return user, nil
}
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	return nil, nil
}
func (r *Repository) PutUser(ctx context.Context, id string, m *model.User) error {
	r.Lock()
	defer r.Unlock()

	r.users[id] = m

	return nil
}
func (r *Repository) DeleteUserByID(ctx context.Context, id string) error {
	return nil
}
