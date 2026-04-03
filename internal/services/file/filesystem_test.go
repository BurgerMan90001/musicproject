package file

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateObject(t *testing.T) {
	t.Parallel()

	temp, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	t.Cleanup(func() {
		os.RemoveAll(temp)
	})

	ctx := context.Background()

	tests := []struct {
		name string

		folder   string
		filename string
		contents []byte

		wantErr error
	}{
		{
			name:     "success",
			folder:   temp,
			filename: "testfile",
			contents: []byte("goona"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFileSystem()
			err := fs.CreateObject(ctx, tt.folder, tt.filename, tt.contents, false, ContentTypeZip)

			// TODO change error asserts to errorIs and refactor error handling
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestGetObject(t *testing.T) {
	t.Parallel()
}

func TestDeleteObject(t *testing.T) {
	t.Parallel()
}
