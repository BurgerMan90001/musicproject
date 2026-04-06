package integration

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
)

// type test struct {
// 	name        string
// 	wantMessage string
// 	wantStatus  int

// 	userId        uuid.UUID
// 	tokenType     string
// 	tokenExpireAt time.Time
// }

func (s *testSuite) TestAuthMiddleware() {
	url := "/v1/protected"

	userId := uuid.Nil

	expired, err := s.jwtService.GenerateToken(userId, auth.TokenAccess, time.Now().Add(-1*time.Hour))
	s.Require().NoError(err)

	invalidType, err := s.jwtService.GenerateToken(userId, auth.TokenRefresh, auth.ExpiresInOneDay)
	s.Require().NoError(err)

	tests := []struct {
		name        string
		wantMessage string
		wantStatus  int
		accessToken string
	}{
		{
			name:        "expired token",
			wantStatus:  http.StatusUnauthorized,
			wantMessage: auth.ErrTokenExpired.Error(),
			accessToken: expired,
		},
		{
			name:        "invalid token type",
			wantMessage: auth.ErrInvalidTokenType.Error(),
			wantStatus:  http.StatusUnauthorized,
			accessToken: invalidType,
		},
		{
			name:        "empty token or no header",
			wantMessage: jwt.ErrTokenMalformed.Error(),
			wantStatus:  http.StatusUnauthorized,
			accessToken: "",
		},
		{
			name:        "invalid token",
			wantStatus:  http.StatusUnauthorized,
			wantMessage: jwt.ErrTokenMalformed.Error(),
			accessToken: "aksidoajfjsofuijrngukaernngaerjgknne",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			req := &request{
				method:      http.MethodGet,
				accessToken: tt.accessToken,
			}
			w := s.newRequest(s.ctx, url, req)

			resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.wantStatus, w.Code)
			s.NotEmpty(resBody["message"])
		})
	}
	valid, err := s.jwtService.GenerateToken(userId, auth.TokenAccess, auth.ExpiresInOneDay)
	s.Require().NoError(err)

	s.Run("success", func() {
		req := &request{
			method:      http.MethodGet,
			accessToken: valid,
		}
		w := s.newRequest(s.ctx, url, req)

		resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
		s.Require().NoError(err)

		s.Equal(http.StatusOK, w.Code)
		s.Empty(resBody["message"])
	})
}

func (s *testSuite) TestRateLimitMiddleware() {

	s.Run("", func() {

	})
}
