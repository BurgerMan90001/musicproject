package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "musicproject.com/gen/mocks"
	"musicproject.com/pkg/model"
)

func TestHandleGet(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		wantBody     string
		wantRepoUser *model.User
		wantStatus   int
		wantErr      error
	}{

		{
			name:       "get user by id",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			wantBody:   `"id":"00000000-0000-0000-0000-000000000000","email":"paul@pol.coom","passwordHash":""}\n`,
			wantRepoUser: &model.User{
				Email: "paul@pol.coom",
			},
		},
		// {
		// 	name:       "invalid method",
		// 	method:     http.MethodPatch,
		// 	wantBody:   "PATCH method not allowed\n",
		// 	wantStatus: http.StatusMethodNotAllowed,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)

			ctx := context.Background()

			id, err := uuid.NewV7()
			if err != nil {
				t.Error(err)
			}
			repoMock.EXPECT().GetUserByID(ctx, id).Return(tt.wantRepoUser, tt.wantErr)

			r, err := http.NewRequestWithContext(ctx, tt.method, fmt.Sprintf("localhost:8081/user?id=%v", id), nil)
			if err != nil {
				t.Error(err)
			}

			w := httptest.NewRecorder()

			handleUser(repoMock).ServeHTTP(w, r)

			res := w.Result()

			var user *model.User
			json.NewDecoder(res.Body).Decode(&user)

			assert.Equal(t, tt.wantStatus, w.Code, tt.name)
			assert.Equal(t, tt.wantRepoUser, user, tt.name)
		})
	}

}
