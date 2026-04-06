package integration

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func (s *testSuite) TestLogin() {
	url := "/v1/auth/login"
	tests := []HandlerTest{
		{
			Name:     "successful login",
			WantCode: http.StatusOK,
			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"email":    "paulcasigay@gmail.com",
					"password": "Dirtycash@123!",
				},
			},
		},
		{
			Name: "incorect password or email",

			WantCode:    http.StatusUnauthorized,
			WantMessage: auth.ErrIncorrectLogin.Error(),

			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"email":    "paulcasigay@gmail.com",
					"password": "Invalidpasswordh@123!",
				},
			},
		},
		{
			Name:     "user not found",
			WantCode: http.StatusNotFound,
			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"email":    "paulcasigayaaaaaaaaa@gmail.com",
					"password": "Dirtycash@123!",
				},
			},
			WantMessage: repository.ErrNotFound.Error(),
		},
		{
			Name: "method not allowed",
			Req: &request{
				method: http.MethodDelete,
			},

			WantCode:    http.StatusMethodNotAllowed,
			WantMessage: handler.ErrInvalidMethod.Error(),
		},
		{
			Name: "empty body",
			Req: &request{
				method: http.MethodPost,
			},

			WantCode:    http.StatusBadRequest,
			WantMessage: handler.ErrInvalidRequestBody.Error(),
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {

			w := s.newRequest(s.ctx, url, tt.Req)

			resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			if tt.WantMessage != "" || resBody["message"] != nil {
				s.Equal(tt.WantMessage, resBody["message"])
			}
			s.Empty(resBody["password"], tt.Name)
			s.Equal(tt.WantCode, w.Code, tt.Name)
		})
	}
}

func (s *testSuite) TestSignup() {
	url := "/v1/auth/signup"

	tests := []HandlerTest{
		{
			Name:     "user already exists",
			WantCode: http.StatusConflict,
			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"email":    "paulcasigay@gmail.com",
					"password": "Dirtycash@123!",
				},
			},

			WantMessage: auth.ErrUserAlreadyExists.Error(),
		},
		{
			Name: "invalid password",

			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"email":    "goop123a@gmail.com",
					"password": "Dirtsh123",
				},
			},
			WantCode:    http.StatusBadRequest,
			WantMessage: auth.ErrInvalidPassword.Error(),
		},
		{
			Name:     "invalid email",
			WantCode: http.StatusBadRequest,
			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"email":    "paulcasigaygmailcom",
					"password": "Dirtycash@123!",
				},
			},

			WantMessage: auth.ErrInvalidEmail.Error(),
		},
		// {
		// 	Name:     "method not allowed",
		// 	Method:   http.MethodDelete,
		// 	WantCode: http.StatusMethodNotAllowed,

		// 	WantMessage: handler.ErrInvalidMethod.Error(),
		// },
		{
			Name: "empty body",
			Req: &request{
				method: http.MethodPost,
			},

			WantCode:    http.StatusBadRequest,
			WantMessage: handler.ErrInvalidRequestBody.Error(),
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, url, tt.Req)

			resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			mes := resBody["message"]
			s.Equal(tt.WantMessage, mes)

			s.Empty(resBody["password"], mes)
			s.Equal(tt.WantCode, w.Code, mes)
		})
	}
	success := HandlerTest{
		Name: "successful signup",
		Req: &request{
			method: http.MethodPost,
			body: map[string]any{
				"email":    "goopay@gmail.com",
				"password": "Dirtycash@123!",
			},
		},
		WantCode: http.StatusCreated,
	}
	s.Run(success.Name, func() {

		w := s.newRequest(s.ctx, url, success.Req)

		resBody, err := jsonutil.ReadJSON[model.User](w.Result().Body)
		s.Require().NoError(err)

		s.Empty(resBody.PasswordHash, success.Name)
		s.Equal(success.Req.body["email"], resBody.Email)
		s.NotEqual(uuid.Nil, resBody.ID)
		s.NotEmpty(resBody.CreatedAt)
	})
}

func (s *testSuite) TestRefresh() {
	url := "/v1/auth/refresh"
	valid, err := s.jwtService.GenerateToken(uuid.Nil, auth.TokenRefresh, auth.ExpiresInOneDay)
	s.Require().NoError(err)

	access, err := s.jwtService.GenerateToken(uuid.Nil, auth.TokenAccess, auth.ExpiresInOneDay)
	s.Require().NoError(err)

	// Gets refresh token from body or cookie
	successTests := []HandlerTest{
		{
			Name:     "success refresh with cookie token",
			WantCode: http.StatusOK,

			Req: &request{
				// Is set in cookie
				refreshToken: valid,
			},
		},
		{
			Name:     "success refresh with response body token",
			WantCode: http.StatusOK,
			Req: &request{
				// Set token in body
				body: map[string]any{
					"refreshToken": valid,
				},
			},
		},
	}

	for _, tt := range successTests {
		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, url, tt.Req)

			resBody := jsonutil.ReadJSONT[map[string]any](s.T(), w.Result().Body)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Empty(resBody["message"])
		})
	}
	failTests := []HandlerTest{
		{
			Name:     "invalid token type",
			WantCode: http.StatusUnauthorized,
			Req: &request{
				refreshToken: access,
			},
			WantMessage: auth.ErrInvalidTokenType.Error(),
		},
		{
			Name:        "empty refresh token string",
			WantMessage: auth.ErrNoRefeshToken.Error(),
			WantCode:    http.StatusUnauthorized,
			Req:         &request{},
		},
	}
	for _, tt := range failTests {
		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, url, tt.Req)

			resBody := jsonutil.ReadJSONT[map[string]any](s.T(), w.Result().Body)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Equal(tt.WantMessage, resBody["message"])
		})
	}
}

/* Oauth tests */
func (s *testSuite) TestHandleOathGoogleLogin() {
	//t := s.T()
	url := "/v1/auth/google/login"
	tests := []HandlerTest{
		{
			Name: "login google oauth",
			Req: &request{
				method: http.MethodGet,
			},

			WantCode: http.StatusOK,
		},
	}
	s.T().Skip()
	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, url, tt.Req)

			userInfo, err := jsonutil.ReadJSON[model.OauthUserInfo](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Empty(userInfo.Email)
		})
	}

}
