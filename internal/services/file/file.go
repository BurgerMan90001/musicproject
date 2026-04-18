package file

import (
	"context"
	"time"
)

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypeZip       = "application/zip"
)

func New(ctx context.Context, T Blobstore, region string) (Blobstore, error) {

	switch T.(type) {
	case *AWSS3:
		return NewS3(ctx, region)
	case *GoogleCloud:
		return NewGoogleCloud(ctx, 30*time.Minute)
	default:
		return NewFileSystem(), nil
	}
}

type Blobstore interface {
	// Creates object in the parent with filename.
	CreateObject(ctx context.Context, bucket, key string,
		contents []byte, cacheble bool, contentType string) error

	// Returns an authenticated url location to upload files at
	CreateObjectUrl(ctx context.Context, bucket, key string,
		cacheble bool) (string, error)

	GetObject(ctx context.Context, bucket, key string) ([]byte, error)
	GetObjectUrl(ctx context.Context, bucket, key string) (string, error)

	DeleteObject(Ctx context.Context, bucket, key string) error
}
