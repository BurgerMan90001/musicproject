package handler

import (
	"io"
	"net/http"
	"testing"

	"movieexample.com/pkg/model"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		body          io.Reader
		wantBody      string
		wantStatus    int
		expectRes     *model.User
		expectRepoRes *model.User
		expectRepoErr error
	}{
		/*
			{
				name:          "put user",
				method:        "PUT",
				url:           "8081/user",
				wantBody:      `{"test":"test"}`,
				wantStatus:    http.StatusOK,
				expectRepoRes: nil,
			},
		*/
		{
			name:       "get user by id",
			method:     "GET",
			wantStatus: http.StatusOK,
			expectRes:  nil,
		},
		/*
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
			/*
				mockController := gomock.NewController(t)
				defer mockController.Finish()

				repoMock := mock_repository.NewMockRepository(mockController)

				ctx := context.Background()
				//repoMock.PutUser(ctx, "id", tt.expectRepoRes)

				r, err := http.NewRequest(tt.method, tt.url, tt.body)
				if err != nil {
					t.Error(err)
				}

				w := httptest.NewRecorder()

				repoMock.EXPECT().GetUserByID(ctx, "id").Return(tt.expectRepoRes, tt.expectRepoErr)

				//authController := auth.New(repoMock, []byte("test_key"))
				userController := user.New(repoMock)

				h := New(nil, userController)

				h.handleUser(w, r)

				res := w.Result()

				var user *model.User
				if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
					t.Errorf("decoding error %v", err)
				}

				if w.Code != tt.wantStatus {
					t.Errorf("wrong status code got: %v wanted: %v", w.Code, tt.wantStatus)
				}

				if user != tt.expectRes {
					t.Errorf("wrong body got: %v wanted: %v", user, tt.expectRes)
				}
			*/
		})
	}
}
