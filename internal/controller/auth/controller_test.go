package auth

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestAuthController(t *testing.T) {
	tests := []struct {
		name          string
		expectRepoErr error
		expectRepoRes string
		wantErr       error
		wantRes       string
	}{
		{
			name:    "dasdasd",
			wantRes: "",
			wantErr: errors.New("awdawd"),
		},
	}
	for _, tt := range tests {
		t.Skip("not implemented yet")
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			// ctx := context.Background()
			// id := "id"

			//assert.Equal(t)
		})
	}
}
