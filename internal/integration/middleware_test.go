package integration

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func (s *testSuite) TestAuthMiddleware() {
	url := "/v1/protected"
	jwtService := auth.NewJWTService(s.cfg.Services.Auth.Jwt)

	tests := []struct {
		name        string
		wantMessage string
		wantStatus  int

		userId        uuid.UUID
		tokenType     string
		tokenExpireAt time.Time
	}{
		{
			name:          "success",
			wantStatus:    http.StatusOK,
			tokenType:     auth.TokenAccess,
			tokenExpireAt: auth.ExpiresInOneDay,
		},
		{
			name:          "expired token",
			wantStatus:    http.StatusUnauthorized,
			wantMessage:   auth.ErrTokenExpired.Error(),
			tokenType:     auth.TokenAccess,
			tokenExpireAt: time.Now().Add(-1 * time.Hour),
		},
		{
			name:          "invalid token type",
			wantMessage:   auth.ErrInvalidTokenType.Error(),
			wantStatus:    http.StatusUnauthorized,
			tokenType:     auth.TokenRefresh,
			tokenExpireAt: auth.ExpiresInOneDay,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tokenString, err := jwtService.GenerateToken(tt.userId, tt.tokenType, tt.tokenExpireAt)
			s.Require().NoError(err)

			w := s.newRequest(s.ctx, "GET", url, nil, tokenString)

			resBody, err := model.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.wantStatus, w.Code)
			if resBody["message"] != nil {
				s.Equal(tt.wantMessage, resBody["message"])
			}
		})
	}

	s.Run("invalid token", func() {
		w := s.newRequest(s.ctx, "GET", url, nil, "aimdlwdnaljngrlgns")

		resBody, err := model.ReadJSON[map[string]any](w.Result().Body)
		s.Require().NoError(err)

		s.Equal(jwt.ErrTokenMalformed.Error(), resBody["message"], "invalid token")
		s.Equal(http.StatusUnauthorized, w.Code, "invalid token")
	})
	s.Run("empty token or no header", func() {
		w := s.newRequest(s.ctx, "GET", url, nil, "")

		resBody, err := model.ReadJSON[map[string]any](w.Result().Body)
		s.Require().NoError(err)

		s.Equal(auth.ErrNoAccessToken.Error(), resBody["message"])
		s.Equal(http.StatusUnauthorized, w.Code)
	})
}

func (s *testSuite) TestRateLimitMiddleware() {

	s.Run("", func() {

	})
}
