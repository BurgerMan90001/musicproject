package upload

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"songsled.com/pkg/model"
)

func TestUploadRequestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		req     *model.SongUploadRequest
		wantErr bool
	}{
		{
			name: "valid upload request",
			req: &model.SongUploadRequest{
				Name:         "asdasd",
				Artists:      []uuid.UUID{},
				Genres:       []uuid.UUID{},
				CreationDate: "2006-25-25",
			},
			wantErr: false,
		},
		{
			name:    "nil request",
			wantErr: true,
		},
		{
			name:    "empty name",
			wantErr: true,
			req: &model.SongUploadRequest{
				Artists:      []uuid.UUID{},
				Genres:       []uuid.UUID{},
				CreationDate: "2006-25-25",
			},
		},
		{
			name:    "empty genre",
			wantErr: true,
			req: &model.SongUploadRequest{
				Name: "asdasd",
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
