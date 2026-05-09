package secrets

import "context"

var _ Manager = (*FileSystem)(nil)

type FileSystem struct {
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (m *FileSystem) Get(ctx context.Context, name string) (string, error) {
	return "", nil
}
func (m *FileSystem) SetEnv(ctx context.Context, name string) error {
	return nil
}
