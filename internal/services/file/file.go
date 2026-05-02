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
		return NewGoogleCloud(ctx, "")
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
		cacheble bool, ttl time.Duration) (string, string, error)

	GetObject(ctx context.Context, bucket, key string) ([]byte, error)
	// Returns an authenticated url location to download files from
	GetObjectUrl(ctx context.Context, bucket, key string, ttl time.Duration) (string, error)

	DeleteObject(Ctx context.Context, bucket, key string) error
}
