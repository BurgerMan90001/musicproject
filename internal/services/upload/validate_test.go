package upload

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"musicproject.com/pkg/model"
)

func TestValidateUploadRequest(t *testing.T) {
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
