package user

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "okapi.com/gen/mocks"
	"okapi.com/pkg/model"
)

func TestUserController(t *testing.T) {
	tests := []struct {
		name          string
		expectRepoErr error
		expectRepoRes *model.User
		wantErr       error
		wantRes       *model.User
	}{
		{
			name:          "success",
			expectRepoRes: &model.User{},
			wantRes:       &model.User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)
			controller := New(repoMock)

			ctx := context.Background()
			id := uuid.New()
			//username := "jane doe"

			repoMock.EXPECT().GetUserByID(ctx, id).Return(tt.expectRepoRes, tt.expectRepoErr)

			res, err := controller.GetUserByID(ctx, id)

			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
