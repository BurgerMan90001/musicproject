package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

type User struct {
	sync.Mutex
	data map[uuid.UUID]*model.User
}

type Repo[T any] struct {
	data map[uuid.UUID]model.User
}

func NewUser() *User {
	return &User{}
}

func (r *User) GetByID(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	r.Lock()
	defer r.Unlock()
	user := r.data[userId]

	return user, nil
}

func (r *User) GetByEmail(ctx context.Context, email string) (*model.User, error) {

	return nil, nil
}
func (r *User) Put(ctx context.Context, user *model.User) (uuid.UUID, error) {
	r.Lock()
	defer r.Unlock()

	if user.ID != uuid.Nil {
		r.data[user.ID] = user
		return user.ID, nil
	}
	userId, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	r.data[userId] = user
	return userId, nil
}

func (r *User) DeleteByID(ctx context.Context, userId uuid.UUID) error {
	delete(r.data, userId)
	return nil
}
