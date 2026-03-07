package memory

import (
	"sync"

	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

type Repository struct {
	sync.RWMutex
	users map[uuid.UUID]*model.User
}

func New() *Repository {
	return &Repository{}
}
