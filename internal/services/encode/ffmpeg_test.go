package encode

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFFmpeg(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	//store := file.NewFileSystem()
	ffmpeg := NewFFmpeg()

	out, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	t.Cleanup(func() {
		os.RemoveAll(out)
	})

	tests := []struct {
		name      string
		inputPath string
		outputDir string
		wantErr   error
	}{
		{
			name:      "success",
			inputPath: "./testdata/8bitBossa.mp3",
			outputDir: out,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ffmpeg.Segment(ctx, tt.inputPath, tt.outputDir)

			assert.Equal(t, tt.wantErr, err, tt.name)
			entries, err := os.ReadDir(out)
			require.NoError(t, err)

			for _, entry := range entries {
				ext := filepath.Ext(entry.Name())

				assertMatchExt(t, ext, ".ts", ".m3u8")
			}
		})

	}
}

// Assert that the file extetion matches
func assertMatchExt(t *testing.T, got string, want ...string) {
	if !slices.Contains(want, got) {
		t.Errorf("want formats: %v got: %v", want, got)
	}
}
