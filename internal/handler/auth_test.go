package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/service/auth"
	authService "musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

func TestLogin(t *testing.T) {
	//url := "/v1/auth"
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
		},
		{
			Name:     "incorect password or email",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
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
			WantMessage: repository.ErrUserNotFound.Error(),
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

			c.userRepo.EXPECT().GetUserByEmail(ctx, tt.Body["email"]).Return(tt.RepoItem, tt.RepoErr).AnyTimes()

			// w := httptest.NewRecorder()

			// body, err := NewRequestBody(tt.Body)
			// if err != nil {
			// 	t.Error(err)
			// }

			//r := httptest.NewRequestWithContext(ctx, tt.Method, "/user", body)

			//r.SetPathValue("id", id.String())
			w, err := newRequest(ctx, tt.Method, tt.URL, tt.Body, false)
			if err != nil {
				t.Error(err)
			}

			//HandleLogin(c.authService).ServeHTTP(w, r)

			resBody, err := model.ReadJSON[model.Response](w.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var data map[string]string
			json.Unmarshal(w.Body.Bytes(), &data)
			_, exists := data["jwt"]

			assert.Equal(t, tt.WantMessage, resBody.Message)
			assert.Equal(t, tt.WantCode, w.Code, tt.Name)

			assert.Equal(t, tt.WantSuccess, exists)

		}))
	}
}
func TestSignup(t *testing.T) {

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
			WantMessage: authService.ErrInvalidEmail.Error(),
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

			c.userRepo.EXPECT().GetUserByEmail(ctx, tt.Body["email"]).Return(tt.RepoItem, tt.RepoErr).AnyTimes()
			c.userRepo.EXPECT().PutUser(ctx, tt.Body["email"], tt.Body["password"]).Return(id, nil).AnyTimes()

			w, err := newRequest(ctx, tt.Method, "/auth/signup", tt.Body, false)
			if err != nil {
				t.Error(err)
			}

			//HandleSignup(c.authService).ServeHTTP(w, r)

			resBody, err := model.ReadJSON[model.Response](w.Result().Body)
			if err != nil {
				t.Error(err)
			}

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

			w, err := newRequest(ctx, tt.Method, tt.URL, tt.Body, false)
			if err != nil {
				t.Error(err)
			}

			// r := httptest.NewRequestWithContext(ctx, tt.Method, "/user", body)

			// HandleOauthLogin(authService).ServeHTTP(w, r)

			//handleOathGoogleRedirect([]byte("test"), oauthCfg).ServeHTTP(w, r)
			// body, err := io.ReadAll(resp.Body)
			// if err != nil {
			// 	t.Error(err)
			// }
			// userInfo, err := ReadJSON[model.GoogleUserInfo](resp.Body)
			// if err != nil {
			// 	t.Errorf("read error: %v", err)
			// }

			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
		})
	}

}
