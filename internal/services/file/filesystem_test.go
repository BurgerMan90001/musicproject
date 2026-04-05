package file

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateObject(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	temp, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	unallowedFolder := "../email"
	testFile := "testfile"

	t.Cleanup(func() {
		os.RemoveAll(temp)
		os.RemoveAll(path.Join(unallowedFolder, testFile))
	})

	fs := NewFileSystem()

	tests := []struct {
		name string

		folder   string
		filename string
		contents []byte
	}{
		// {
		// 	name:     "unallowed folder",
		// 	folder:   unallowedFolder,
		// 	filename: testFile,
		// },
		{
			name:     "invalid folder",
			folder:   "../../somntiosrtart",
			filename: "testfile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			err := fs.CreateObject(ctx, tt.folder, tt.filename, tt.contents, false, ContentTypeZip)

			// TODO change error asserts to errorIs and refactor error handling
			assert.Error(t, err, tt.name)
		})
	}
	t.Run("success", func(t *testing.T) {
		err := fs.CreateObject(ctx, temp, testFile, []byte("goona"), false, ContentTypeZip)
		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(temp, testFile))
		require.NoError(t, err)
		assert.Equal(t, content, []byte("goona"))
	})
}

func TestGetObject(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name     string
		folder   string
		filename string
	}{
		{
			name: "success",
		},
		{
			name: "invalid folder",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

	fs := NewFileSystem()
	fs.GetObject(ctx, "", "")
}

func TestDeleteObject(t *testing.T) {
	t.Parallel()
}
