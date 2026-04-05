package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

type User struct {
	mu   sync.Mutex
	data map[uuid.UUID]*model.User
}

func NewUser() *User {

	return &User{data: make(map[uuid.UUID]*model.User, 10)}
}

func (r *User) GetByID(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.data[userId]
	if !exists {
		return nil, repository.ErrNotFound
	}

	return user, nil
}

func (r *User) GetByEmail(ctx context.Context, email string) (*model.User, error) {

	return nil, nil
}
func (r *User) Put(ctx context.Context, user *model.User) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// if user.ID != uuid.Nil {
	// 	r.data[user.ID] = user
	// 	return user.ID, nil
	// }
	// userId, err := uuid.NewV7()
	// if err != nil {
	// 	return uuid.Nil, err
	// }
	r.data[user.ID] = user
	return user.ID, nil
}

func (r *User) DeleteByID(ctx context.Context, userId uuid.UUID) error {
	delete(r.data, userId)
	return nil
}
