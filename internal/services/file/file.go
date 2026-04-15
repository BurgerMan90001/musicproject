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
	CreateObject(ctx context.Context, parent, filename string,
		contents []byte, cacheble bool, contentType string) error

	// Returns an authenticated url location to upload files at
	CreateObjectUrl(ctx context.Context, parent, filename string,
		cacheble bool) (string, error)

	GetObject(ctx context.Context, parent, name string) ([]byte, error)
	DeleteObject(Ctx context.Context, parrent, name string) error
}
