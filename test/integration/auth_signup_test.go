package integration

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func (s *testSuite) TestSignup() {
	url := "/v1/auth/signup"

	tests := []handlerTest{
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
		{
			Name: "empty body",
			Req: &request{
				method: http.MethodPost,
			},

			WantCode:    http.StatusBadRequest,
			WantMessage: "Invalid request body",
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)

			resBody, err := jsonutil.ReadJson[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			mes := resBody["message"]
			s.Equal(tt.WantMessage, mes)

			s.Empty(resBody["password"], mes)
			s.Equal(tt.WantCode, w.Code, mes)
		})
	}
	success := handlerTest{
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

		w := s.newRequest(url, success.Req)

		resBody, err := jsonutil.ReadJson[model.User](w.Result().Body)
		s.Require().NoError(err)

		s.Empty(resBody.PasswordHash, success.Name)
		s.Equal(success.Req.body["email"], resBody.Email)
		s.NotEqual(uuid.Nil, resBody.ID)
		s.NotEmpty(resBody.CreatedAt)
	})
}
