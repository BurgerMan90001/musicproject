package file

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var _ Blobstore = (*FileSystem)(nil)

type FileSystem struct{}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (s *FileSystem) CreateObject(ctx context.Context, folder, filename string,
	contents []byte, cacheble bool, contentType string) error {
	if filename == "" {
		return errors.New("Create object: filename is empty")
	}
	path := filepath.Join(folder, filename)

	if err := os.WriteFile(path, contents, 0o600); err != nil {
		return fmt.Errorf("Create object: %w", err)
	}
	return nil
}
func (s *FileSystem) CreateObjectUrl(ctx context.Context, folder, filename string,
	cacheble bool) (string, error) {
	if filename == "" {
		return "", errors.New("Create object: filename is empty")
	}
	path := filepath.Join(folder, filename)

	// if err := os.WriteFile(path, contents, 0o600); err != nil {
	// 	return "", fmt.Errorf("Create object: %w", err)
	// }
	return path, nil
}
func (s *FileSystem) GetObject(ctx context.Context, folder string, filename string) ([]byte, error) {
	path := filepath.Join(folder, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}

	return data, nil
}

func (s *FileSystem) GetObjectUrl(ctx context.Context, bucket, key string) (string, error) {
	return "", nil
}
func (s *FileSystem) DeleteObject(Ctx context.Context, folder string, filename string) error {
	path := filepath.Join(folder, filename)
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
