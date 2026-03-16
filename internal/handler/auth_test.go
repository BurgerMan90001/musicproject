package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "musicproject.com/gen/mocks"
	"musicproject.com/internal/repository"
	authService "musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

func TestLogin(t *testing.T) {
	jwtKey := []byte("test")
	tests := []HandlerTest{
		{
			Name:     "successful login",
			Method:   http.MethodGet,
			WantCode: http.StatusOK,
			Body: map[string]any{
				"email": "paul@pol.com",
			},
		},
	}
	t.Skip()
	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)

			ctx := context.Background()

			id, err := uuid.NewV7()
			if err != nil {
				t.Error(err)
			}

			repoMock.EXPECT().GetUserByID(ctx, id).Return(tt.RepoItem, tt.RepoErr)

			w := httptest.NewRecorder()

			body, err := NewRequestBody(tt.Body)
			if err != nil {
				t.Error(err)
			}

			r := httptest.NewRequestWithContext(ctx, tt.Method, "/user", body)
			r.SetPathValue("id", id.String())
			HandleLogin(jwtKey, repoMock).ServeHTTP(w, r)
			//HandleUser(repoMock).ServeHTTP(w, r)

			// var user *model.User
			// if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
			// 	t.Error(err)
			// }
			//res := w.Result()
			// resBody, err := io.ReadAll(res.Body)
			// if err != nil {
			// 	t.Error(err)
			// }

			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
			//assert.Equal(t, tt.WantRes, string(resBody))

			//assert.Equal(t, tt.)
			//assert.Equal(t, tt.wantUser, user, tt.name)
			//repoMock.EXPECT().PutUser(ctx, id, user, )
		})
	}
}
func TestSignup(t *testing.T) {
	id, err := uuid.NewV7()
	if err != nil {
		t.Error(err)
	}
	jwtKey := []byte("test")
	tests := []HandlerTest{
		{
			Name:     "successful signup",
			Method:   http.MethodPost,
			WantCode: http.StatusOK,
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},
			// WantData: map[string]any{
			// 	"id":    id.String(),
			// 	"email": "paulcasigay@gmail.com",
			// },
			WantStatus: StatusSucess,
			RepoErr:    repository.ErrNotFound,
		},
		{
			Name:     "user already exists",
			Method:   http.MethodPost,
			WantCode: http.StatusConflict,
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},

			WantStatus:  StatusError,
			WantMessage: ErrUserAlreadyExists.Error(),

			RepoItem: &model.User{
				ID:           id,
				Email:        "paulcasigay@gmail.com",
				PasswordHash: "Dirtycash@123!",
			},
		},
		{
			Name:     "invalid password",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtsh123",
			},
			WantStatus:  StatusError,
			WantMessage: authService.ErrInvalidPassword.Error(),
		},
		{
			Name:     "invalid email",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,

			Body: map[string]any{
				"email":    "paulcasigaygmailcom",
				"password": "Dirtycash@123!",
			},
			WantStatus:  StatusError,
			WantMessage: authService.ErrInvalidEmail.Error(),
		},
		{
			Name:     "method not allowed",
			Method:   http.MethodDelete,
			WantCode: http.StatusMethodNotAllowed,

			WantStatus:  StatusError,
			WantMessage: ErrInvalidMethod.Error(),
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)

			ctx := context.Background()

			repoMock.EXPECT().GetUserByEmail(ctx, tt.Body["email"]).Return(tt.RepoItem, tt.RepoErr).AnyTimes()

			repoMock.EXPECT().PutUser(ctx, tt.Body["email"], tt.Body["password"]).Return(id, nil).AnyTimes()

			w := httptest.NewRecorder()

			body, err := NewRequestBody(tt.Body)
			if err != nil {
				t.Error(err)
			}

			r := httptest.NewRequestWithContext(ctx, tt.Method, "/user", body)
			r.Header.Set("Content-Type", "application/json")

			HandleSignup(jwtKey, repoMock).ServeHTTP(w, r)

			t1, err := MarshalJSON(tt.WantStatus, tt.WantData, tt.WantCode, tt.WantMessage)
			if err != nil {
				t.Error(err)
			}

			t2, err := io.ReadAll(w.Body)
			if err != nil {
				t.Error(err)
			}

			assert.JSONEq(t, string(t1), string(t2), tt.Name)
			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
		})
	}
}
