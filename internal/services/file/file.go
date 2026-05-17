package file

import (
	"context"
	"time"

	"songsled.com/internal/config"
)

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypeZip       = "application/zip"
)

// public: The public endpoint where content is served from
func New(ctx context.Context, T Blobstore, cfg config.File) (Blobstore, error) {

	switch T.(type) {
	case *AWSS3:
		return NewS3(ctx, cfg.Endpoint, cfg.Region, cfg.Public, "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY")
	case *GoogleCloud:
		return NewGoogleCloud(ctx, cfg.Endpoint, "GOOGLE_ACCESS_ID")
		
	default:
		return NewFileSystem(), nil
	}
}

type Blobstore interface {
	// Creates object in the parent with filename.
	CreateObject(ctx context.Context, bucket, key string,
		contents []byte, cacheble bool, contentType string) error

	// Returns a presigned url to upload files at
	// and the url where the object can be download from
	CreateObjectUrl(ctx context.Context, bucket, key string,
		cacheble bool, ttl time.Duration, contentType string) (string, string, error)

	GetObject(ctx context.Context, bucket, key string) ([]byte, error)
	// Returns an authenticated url location to download files from
	GetObjectUrl(ctx context.Context, bucket, key string, ttl time.Duration) (string, error)

	DeleteObject(Ctx context.Context, bucket, key string) error
}
