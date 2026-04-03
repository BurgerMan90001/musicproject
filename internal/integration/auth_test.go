package integration

import (
	"net/http"

	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func (s *testSuite) TestLogin() {
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
		},
		{
			Name:        "incorect password or email",
			Method:      http.MethodPost,
			WantCode:    http.StatusUnauthorized,
			WantMessage: auth.ErrIncorrectLogin.Error(),
			Body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Invalidpasswordh@123!",
			},
		},
		{
			Name:     "user not found",
			Method:   http.MethodPost,
			WantCode: http.StatusNotFound,
			Body: map[string]any{
				"email":    "paulcasigayaaaaaaaaa@gmail.com",
				"password": "Dirtycash@123!",
			},
			WantMessage: repository.ErrNotFound.Error(),
		},
		{
			Name:     "method not allowed",
			Method:   http.MethodDelete,
			WantCode: http.StatusMethodNotAllowed,

			WantMessage: handler.ErrInvalidMethod.Error(),
		},
		{
			Name:     "empty body",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,

			WantMessage: handler.ErrInvalidRequestBody.Error(),
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, tt.Method, url, tt.Body, "")

			resBody, err := model.ReadJSON[map[string]any](w.Result().Body)
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
			Name:     "successful signup",
			Method:   http.MethodPost,
			WantCode: http.StatusCreated,
			Body: map[string]any{
				"email":    "goopay@gmail.com",
				"password": "Dirtycash@123!",
			},
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
		},
		{
			Name:     "invalid password",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,
			Body: map[string]any{
				"email":    "goop123a@gmail.com",
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

			WantMessage: handler.ErrInvalidMethod.Error(),
		},
		{
			Name:     "empty body",
			Method:   http.MethodPost,
			WantCode: http.StatusBadRequest,

			WantMessage: handler.ErrInvalidRequestBody.Error(),
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, tt.Method, url, tt.Body, "")

			resBody, err := model.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			if tt.WantMessage != "" || resBody["message"] != nil {
				s.Equal(tt.WantMessage, resBody["message"], tt.Name)
			}

			s.Empty(resBody["password"], tt.Name)
			s.Equal(tt.WantCode, w.Code, tt.Name)

		})
	}
}

/* Oauth tests */
func (s *testSuite) TestHandleOathGoogleLogin() {
	//t := s.T()
	url := "/v1/auth/google/login"
	tests := []HandlerTest{
		{
			Name:     "login google oauth",
			Method:   "GET",
			WantCode: http.StatusOK,
			WantData: model.TokenPair{},
		},
	}
	s.T().Skip()
	for _, tt := range tests {

		s.Run(tt.Name, func() {
			w := s.newRequest(s.ctx, tt.Method, url, tt.Body, "")

			userInfo, err := model.ReadJSON[model.OauthUserInfo](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Empty(userInfo.Email)
		})
	}

}
