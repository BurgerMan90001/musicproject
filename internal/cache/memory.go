package cache

import (
	"sync"
	"time"
)

// var _ Cache = (*Memory)(nil)
const size = 15

type Memory[T any] struct {
	data map[string]item[T]
	mu   sync.RWMutex
}

type item[T any] struct {
	object    T
	expiresAt int64
}

func (i *item[T]) expired() bool {
	return i.expiresAt < time.Now().UnixNano()
}
func NewMemory[T any]() *Memory[T] {
	return &Memory[T]{
		data: make(map[string]item[T]),
	}
}

func (c *Memory[T]) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.data)
}

func (c *Memory[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]item[T], size)
}
