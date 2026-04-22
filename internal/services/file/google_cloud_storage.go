package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

var _ Blobstore = (*GoogleCloud)(nil)

type GoogleCloud struct {
	client   *storage.Client
	accessId string
}

func NewGoogleCloud(ctx context.Context) (*GoogleCloud, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("New google cloud store: %w", err)
	}

	accessId := os.Getenv("GOOGLE_ACCESS_ID")

	return &GoogleCloud{
		client:   client,
		accessId: accessId,
	}, nil
}

func (s *GoogleCloud) CreateObject(ctx context.Context, bucket, key string,
	contents []byte, cacheble bool, contentType string) error {
	cacheControl := "public, max-age=86400"
	if !cacheble {
		cacheControl = "no-cache, max-age=0"
	}

	w := s.client.Bucket(bucket).Object(key).NewWriter(ctx)
	w.CacheControl = cacheControl
	if contentType != "" {
		w.ContentType = contentType
	}

	if _, err := w.Write(contents); err != nil {
		return fmt.Errorf("Create object write: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("Create object writer close: %w", err)
	}

	return nil
}

func (s *GoogleCloud) CreateObjectUrl(ctx context.Context, bucket, key string,
	cacheble bool, ttl time.Duration) (string, error) {
	opts := &storage.SignedURLOptions{
		GoogleAccessID: s.accessId,
		Method:         "PUT",
		Expires:        time.Now().Add(ttl),
		Scheme:         storage.SigningSchemeV4,
	}
	//s.client.Bucket()

	return s.client.Bucket(bucket).SignedURL(key, opts)
}

func (s *GoogleCloud) GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	r, err := s.client.Bucket(bucket).Object(key).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Get object new reader: %w", err)
	}
	defer r.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return nil, fmt.Errorf("Download object bytes: %w", err)
	}
	return b.Bytes(), nil
}
func (s *GoogleCloud) GetObjectUrl(ctx context.Context,
	bucket, key string, ttl time.Duration) (string, error) {
	opts := &storage.SignedURLOptions{
		GoogleAccessID: s.accessId,
		Method:         "GET",
		Expires:        time.Now().Add(ttl),
	}

	return s.client.Bucket(bucket).SignedURL(key, opts)
}
func (s *GoogleCloud) DeleteObject(ctx context.Context, bucket, key string) error {
	return s.client.Bucket(bucket).Object(key).Delete(ctx)
}
