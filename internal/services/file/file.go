package file

import (
	"context"
)

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypeZip       = "application/zip"
)

type Blobstore interface {
	// Creates object in the parent with filename.
	CreateObject(ctx context.Context, parent, name string,
		contents []byte, cacheble bool, contentType string) error
	GetObject(ctx context.Context, parent string, name string) ([]byte, error)
	DeleteObject(Ctx context.Context, parrent string, name string) error
}
