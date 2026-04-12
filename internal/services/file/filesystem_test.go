package file

import (
	"os"
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

	t.Cleanup(func() {
		os.RemoveAll(temp)
	})

	fs := NewFileSystem()

	tests := []struct {
		name string

		folder   string
		filename string
		contents []byte

		wantErr bool
	}{
		{
			name:     "success",
			folder:   temp,
			filename: "test.txt",
			contents: []byte("goona"),
		},
		{
			name:     "invalid folder",
			folder:   "../../somntiosrtart",
			filename: "testfile",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			err := fs.CreateObject(ctx, tt.folder, tt.filename, tt.contents, false, ContentTypeZip)

			// TODO change error asserts to errorIs and refactor error handling
			if tt.wantErr {
				assert.Error(t, err, tt.name)
				assert.NoFileExists(t, filepath.Join(tt.folder, tt.filename))
			} else {
				assert.NoError(t, err)
				content, err := os.ReadFile(filepath.Join(tt.folder, tt.filename))
				require.NoError(t, err)

				assert.Equal(t, string(tt.contents), string(content))
			}
		})
	}
}

func TestGetObject(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	temp, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	t.Cleanup(func() {
		os.RemoveAll(temp)
	})

	fs := NewFileSystem()

	tests := []struct {
		name     string
		folder   string
		filename string
		wantData []byte
		wantErr  bool
	}{
		{
			name:     "success",
			folder:   temp,
			filename: "hello.txt",

			wantData: []byte("hello world!"),
			wantErr:  false,
		},
		{
			name:     "invalid folder",
			filename: "hello.txt",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fs.CreateObject(ctx, temp, "hello.txt", []byte("hello world!"), false, "")
			require.NoError(t, err)

			data, err := fs.GetObject(ctx, tt.folder, tt.filename)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				require.NoError(t, err)
				assert.Equal(t, string(tt.wantData), string(data))
			}

		})
	}

}

func TestDeleteObject(t *testing.T) {
	t.Parallel()
}
