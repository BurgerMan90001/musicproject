package upload

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"songsled.com/pkg/model"
)

func TestUploadRequestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		req     *model.UploadSongRequest
		wantErr bool
	}{
		{
			name: "valid upload request",
			req: &model.UploadSongRequest{
				Name:     "asdasd",
				Genre:    "asdasd",
				Image:    "asdasd",
				Filename: "awwww",
			},
		},
		{
			name:    "nil request",
			wantErr: true,
		},
		{
			name:    "empty name",
			wantErr: true,
			req: &model.UploadSongRequest{
				Genre:    "asdasd",
				Image:    "asdasd",
				Filename: "awwww",
			},
		},
		{
			name:    "empty genre",
			wantErr: true,
			req: &model.UploadSongRequest{
				Name:     "asdasd",
				Image:    "asdasd",
				Filename: "awwww",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := validateUploadRequest(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
