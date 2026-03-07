package memory

import (
	"context"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
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
func (r *Repository) PutUser(ctx context.Context, id uuid.UUID, m *model.User) error {
	r.Lock()
	defer r.Unlock()

	r.users[id] = m

	return nil
}
func (r *Repository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return nil
}
