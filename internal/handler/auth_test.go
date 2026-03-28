package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func TestLogin(t *testing.T) {
	url := "/v1/auth/login"
	tests := []HandlerTest{
		{
			Name:     "successful login",
			Method:   http.MethodPost,
			WantCode: http.StatusOK,

			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},

			RepoItem: &model.User{
				ID:           id,
				Email:        "paulcasigay@gmail.com",
				PasswordHash: "Dirtycash@123!",
			},
			//WantData: true,
		},
		{
			Name:        "incorect password or email",
			Method:      http.MethodPost,
			WantCode:    http.StatusUnauthorized,
			WantMessage: "incorrect password or email",
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},
			RepoItem: &model.User{
				ID:           id,
				Email:        "paulcasigay@gmail.com",
				PasswordHash: "Invalidpasswordh@123!",
			},
		},
		{
			Name:     "user not found",
			Method:   http.MethodPost,
			WantCode: http.StatusNotFound,
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},
			RepoErr:     repository.ErrNotFound,
			WantMessage: repository.ErrNotFound.Error(),
		},
		{
			Name:     "method not allowed",
			Method:   http.MethodDelete,
			WantCode: http.StatusMethodNotAllowed,

			WantMessage: ErrInvalidMethod.Error(),
		},
	}
	//t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, testCase(func(t *testing.T, c *testContext) {
			user, ok := tt.RepoItem.(*model.User)
			if ok {
				passwordHash, err := auth.HashPassword(user.PasswordHash)
				if err != nil {
					t.Error(err)
				}
				user.PasswordHash = passwordHash
			}
			c.repo.EXPECT().GetUserByEmail(ctx, tt.Body["email"]).Return(user, tt.RepoErr).AnyTimes()

			w, err := newRequest(ctx, tt.Method, url, tt.Body, "")
			assert.NoError(t.)

			resBody, err := model.ReadJSON[model.Response](w.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var data map[string]string
			json.Unmarshal(w.Body.Bytes(), &data)
			_, exists := data["jwt"]

			assert.Equal(t, tt.WantMessage, resBody.Message)
			assert.Equal(t, tt.WantCode, w.Code, tt.Name)

			assert.Equal(t, tt.WantData, exists)

		}))
	}
}
func TestSignup(t *testing.T) {
	url := "/v1/auth/signup"

	// booleans will default to false if not assigned
	tests := []HandlerTest{
		{
			Name:     "successful signup",
			Method:   http.MethodPost,
			WantCode: http.StatusOK,
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},

			WantSuccess: true,

			// The user is new and has not previously signed up
			RepoErr: repository.ErrNotFound,
		},
		{
			Name:     "user already exists",
			Method:   http.MethodPost,
			WantCode: http.StatusConflict,

			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},

			WantMessage: auth.ErrUserAlreadyExists.Error(),

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
			WantMessage: auth.ErrInvalidPassword.Error(),
		},
		{
			Name:     "invalid email",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,

			Body: map[string]any{
				"email":    "paulcasigaygmailcom",
				"password": "Dirtycash@123!",
			},
			WantMessage: auth.ErrInvalidEmail.Error(),
		},
		{
			Name:     "method not allowed",
			Method:   http.MethodDelete,
			WantCode: http.StatusMethodNotAllowed,

			WantMessage: ErrInvalidMethod.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, testCase(func(t *testing.T, c *testContext) {
			c.repo.EXPECT().GetUserByEmail(ctx, tt.Body["email"]).Return(tt.RepoItem, tt.RepoErr).AnyTimes()
			c.repo.EXPECT().PutUser(ctx, tt.Body["email"], gomock.Any()).Return(id, nil).AnyTimes()

			w, err := newRequest(ctx, tt.Method, url, tt.Body, "")
			assert.NoError(t, err)

			resBody, err := model.ReadJSON[model.Response](w.Result().Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.WantMessage, resBody.Message)
			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
		}))
	}
}

/* Oauth tests */
func TestHandleOathGoogleLogin(t *testing.T) {
	tests := []HandlerTest{
		{
			Name:     "login google oauth",
			Method:   "GET",
			WantCode: http.StatusOK,
			WantData: model.TokenPair{},
		},
	}
	t.Skip()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			w, err := newRequest(ctx, tt.Method, tt.URL, tt.Body, "")
			if err != nil {
				t.Error(err)
			}

			userInfo, err := model.ReadJSON[model.OauthUserInfo](w.Result().Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
			assert.Equal(t, "", userInfo.Email)
		})
	}

}
