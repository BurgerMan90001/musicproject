package file

import (
	"context"
)

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypeZip       = "application/zip"
)

type Blobstore interface {
	CreateObject(ctx context.Context, parent string, name string,
		contents []byte, cacheble bool, contentType string) error
	GetObject(ctx context.Context, parent string, name string) ([]byte, error)
	DeleteObject(Ctx context.Context, parrent string, name string) error
}
