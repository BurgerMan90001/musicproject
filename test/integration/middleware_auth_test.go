package integration

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func (s *testSuite) TestAuthMiddleware() {
	url := "/v1/protected"

	userId := uuid.Nil

	expireJwt, err := auth.NewJWTService(
		s.cfg.Services.Auth.Jwt, "JWT_ACCESS_KEY",
		model.TokenAccess, -1*time.Hour)

	expired, err := expireJwt.GenerateToken(userId, nil)
	s.Require().NoError(err)

	invalidType, err := s.jwtRefresh.GenerateToken(userId, nil)
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
			w := s.newRequest(url, req)

			resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.wantStatus, w.Code, tt.name)
			s.NotEmpty(resBody["message"], tt.name)
		})
	}
	valid, err := s.jwtAccess.GenerateToken(userId, nil)
	s.Require().NoError(err)

	s.Run("success", func() {
		req := &request{
			method:      http.MethodGet,
			accessToken: valid,
		}
		w := s.newRequest(url, req)

		resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
		s.Require().NoError(err)

		s.Equal(http.StatusOK, w.Code)
		s.Empty(resBody["message"])
	})
}
