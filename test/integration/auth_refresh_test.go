package integration

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/services/auth"
)

func (s *testSuite) TestRefreshSuccess() {
	url := "/v1/auth/refresh"
	valid, err := s.jwtRefresh.GenerateToken(uuid.Nil, nil)
	s.Require().NoError(err)
	// Gets refresh token from body or cookie
	tests := []HandlerTest{
		{
			Name:     "successful refresh with cookie token",
			WantCode: http.StatusOK,

			Req: &request{
				// Set token in cookie
				refreshToken: valid,
			},
		},
		// {
		// 	Name:     "successful refresh with request body token",
		// 	WantCode: http.StatusOK,
		// 	Req: &request{
		// 		// Set token in body
		// 		body: map[string]any{
		// 			"refreshToken": valid,
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)

			//resBody := jsonutil.ReadJSONT[map[string]any](s.T(), w.Result().Body)

			s.Equal(tt.WantCode, w.Code, tt.Name)

			//s.Empty(resBody)
			// Make request with old refresh token
			w2 := s.newRequest(url, tt.Req)
			// Old refresh token is revoked
			s.Equal(http.StatusUnauthorized, w2.Code, tt.Name)
		})
	}

}
func (s *testSuite) TestRefresh() {
	url := "/v1/auth/refresh"

	access, err := s.jwtAccess.GenerateToken(uuid.Nil, nil)
	s.Require().NoError(err)

	tests := []HandlerTest{
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
	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)

			//resBody := jsonutil.ReadJSONT[model.ErrorResponse](s.T(), w.Result().Body)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			//s.Equal(tt.WantMessage, resBody.Message)
		})
	}
}
