package file

import (
	"context"
	"sync"
)

var _ Blobstore = (*Memory)(nil)

type Memory struct {
	mu   sync.Mutex
	data map[string][]byte
}

func NewMemory() *Memory {
	return &Memory{
		data: make(map[string][]byte),
	}
}

func (s *Memory) CreateObject(ctx context.Context, parent string, name string,
	contents []byte, cacheble bool, contentType string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil
}
func (s *Memory) CreateObjectUrl(ctx context.Context, parent, name string,
	cacheble bool) (string, error) {
	return "", nil
}
func (s *Memory) GetObject(ctx context.Context, parent string, name string) ([]byte, error) {
	return nil, nil

}
func (s *Memory) DeleteObject(Ctx context.Context, parrent string, name string) error {
	return nil
}
