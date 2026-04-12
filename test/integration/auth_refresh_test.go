package integration

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
)

func (s *testSuite) TestRefresh() {
	url := "/v1/auth/refresh"

	// jwtAccess, err := auth.NewJWTService(s.cfg.Services.Auth.Jwt, "JWT_ACCESS_KEY", model.TokenAccess, time.Hour)
	// s.Require().NoError(err)

	valid, err := s.jwtRefresh.GenerateToken(uuid.Nil, nil)
	s.Require().NoError(err)

	access, err := s.jwtAccess.GenerateToken(uuid.Nil, nil)
	s.Require().NoError(err)

	// Gets refresh token from body or cookie
	successTests := []HandlerTest{
		{
			Name:     "successful refresh with cookie token",
			WantCode: http.StatusOK,

			Req: &request{
				// Set token in cookie
				refreshToken: valid,
			},
		},
		{
			Name:     "successful refresh with request body token",
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
			w := s.newRequest( url, tt.Req)

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
			WantMessage: auth.ErrNoRefeshToken.Error(),
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
			w := s.newRequest(url, tt.Req)

			resBody := jsonutil.ReadJSONT[map[string]any](s.T(), w.Result().Body)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Equal(tt.WantMessage, resBody["message"])
		})
	}
}
