package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
	mock_repository "movieexample.com/gen/mocks"
	"movieexample.com/internal/controller/user"
	"movieexample.com/pkg/model"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		url           string
		body          io.Reader
		wantBody      string
		wantStatus    int
		expectRepoRes *model.User
		expectRepoErr error
	}{
		{
			name:          "put user",
			method:        "PUT",
			url:           "8081/user",
			wantBody:      `{"test":"test"}`,
			wantStatus:    http.StatusOK,
			expectRepoRes: nil,
		},
		{
			name:          "get user by id",
			method:        "GET",
			url:           "/user?id=id",
			wantStatus:    http.StatusOK,
			expectRepoRes: &model.User{ID: "", Username: ""},
		},
		{
			name:       "invalid method",
			method:     http.MethodDelete,
			url:        "/user?id=id",
			wantBody:   "invalid request method",
			wantStatus: http.StatusMethodNotAllowed,
		},
		/*
			{
				name:       "signup success",
				method:     "POST",
				url:        "localhost:8081/auth/signup",
				expect:     `{"test":"test"}`,
				wantStatus: http.StatusOK,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)

			r, err := http.NewRequest(tt.method, tt.url, tt.body)
			if err != nil {
				t.Error(err)
			}

			w := httptest.NewRecorder()

			ctx := context.Background()

			repoMock.EXPECT().GetUserByID(ctx, "id").Return(tt.expectRepoRes, tt.expectRepoErr)

			//authController := auth.New(repoMock, []byte("test_key"))
			userController := user.New(repoMock)

			h := New(nil, userController)
			h.handleUser(w, r)

			res := w.Result()
			data, err := io.ReadAll(res.Body)

			if w.Code != tt.wantStatus {
				t.Errorf("wrong status code got: %v wanted: %v", w.Code, tt.wantStatus)
			}

			if string(data) != tt.wantBody {
				t.Errorf("wrong body got: %v wanted: %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}
