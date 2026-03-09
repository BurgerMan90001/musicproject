package user

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "musicproject.com/gen/mocks"
	"musicproject.com/pkg/model"
)

func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name string

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

			id, err := uuid.NewV7()
			if err != nil {
				t.Error(err)
			}
			ctx := context.Background()

			repoMock.EXPECT().GetUserByID(ctx, id).Return(tt.expectRepoRes, tt.expectRepoErr)

			res, err := controller.GetUserByID(ctx, id)

			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
func TestPutUser(t *testing.T) {
	tests := []struct {
		name          string
		user          *model.User
		wantUser      *model.User
		expectRepoErr error
		wantErr       error
	}{
		{
			name: "success",
			user: &model.User{
				Email:        "paul@doe.com",
				PasswordHash: "test123a",
			},
			wantErr: nil,
		},
		{
			name: "empty user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)
			controller := New(repoMock)

			ctx := context.Background()

			id := uuid.Nil

			repoMock.EXPECT().PutUser(ctx, id, tt.user).Return(tt.expectRepoErr)

			err := controller.PutUser(ctx, id, tt.user.Email, tt.user.PasswordHash)

			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
