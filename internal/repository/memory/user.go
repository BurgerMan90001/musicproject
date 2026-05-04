package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"songsled.com/internal/repository"
	"songsled.com/pkg/model"
)

type User struct {
	mu     sync.RWMutex
	data   map[uuid.UUID]*model.User
	emails map[string]uuid.UUID
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
	r.mu.Lock()
	defer r.mu.Unlock()

	userId, exists := r.emails[email]
	if !exists {
		return nil, repository.ErrNotFound
	}
	user, err := r.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *User) Put(ctx context.Context, user *model.User) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	userId, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	r.data[userId] = user
	// r.emails[user.Email] = userId
	return user.ID, nil
}

func (r *User) DeleteByID(ctx context.Context, userId uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, userId)
	return nil
}
