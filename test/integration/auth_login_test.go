package integration

import (
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func (s *testSuite) TestLogin() {
	url := "/v1/auth/login"
	failTests := []handlerTest{
		{
			Name:        "incorect password or email",
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
			WantMessage: "Method not allowed",
		},
		{
			Name: "empty body",
			Req: &request{
				method: http.MethodPost,
			},

			WantCode:    http.StatusBadRequest,
			WantMessage: "Invalid request body",
		},
	}

	for _, tt := range failTests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)

			resBody, err := jsonutil.ReadJson[model.Error](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.WantMessage, resBody.Message)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Equal(tt.WantCode, resBody.Code, tt.Name)
		})
	}
	tt := handlerTest{
		Name:     "successful login",
		WantCode: http.StatusOK,
		Req: &request{
			method: http.MethodPost,
			body: map[string]any{
				"email":    "paulcasigay@gmail.com",
				"password": "Dirtycash@123!",
			},
		},
	}
	s.Run(tt.Name, func() {
		w := s.newRequest(url, tt.Req)

		user, err := jsonutil.ReadJson[model.User](w.Result().Body)
		s.Require().NoError(err)

		s.Equal("paulcasigay@gmail.com", user.Email)
		s.Empty(user.PasswordHash)

		s.NotEmpty(w.Result().Cookies())
		// s.Contains(w.Result().Cookies(), auth.TokenAccess)
		// s.Contains(w.Result().Cookies(), auth.TokenRefresh)

		s.Equal(tt.WantCode, w.Code, tt.Name)
	})
}
