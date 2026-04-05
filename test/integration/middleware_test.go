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

	jwtService, err := auth.NewJWTService(s.ctx, s.sm)
	s.Require().NoError(err)

	userId := uuid.Nil

	expired, err := jwtService.GenerateToken(userId, auth.TokenAccess, time.Now().Add(-1*time.Hour))
	s.Require().NoError(err)

	invalidType, err := jwtService.GenerateToken(userId, auth.TokenRefresh, auth.ExpiresInOneDay)
	s.Require().NoError(err)

	tests := []struct {
		name        string
		wantMessage string
		wantStatus  int
		tokenString string
	}{
		{
			name:        "expired token",
			wantStatus:  http.StatusUnauthorized,
			wantMessage: auth.ErrTokenExpired.Error(),
			tokenString: expired,
		},
		{
			name:        "invalid token type",
			wantMessage: auth.ErrInvalidTokenType.Error(),
			wantStatus:  http.StatusUnauthorized,
			tokenString: invalidType,
		},
		{
			name:        "empty token or no header",
			wantMessage: jwt.ErrTokenMalformed.Error(),
			wantStatus:  http.StatusUnauthorized,
			tokenString: "",
		},
		{
			name:        "invalid token",
			wantStatus:  http.StatusUnauthorized,
			wantMessage: jwt.ErrTokenMalformed.Error(),
			tokenString: "aksidoajfjsofuijrngukaernngaerjgknne",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			w := s.newRequest(s.ctx, "GET", url, nil, tt.tokenString)

			resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.wantStatus, w.Code)
			s.NotEmpty(resBody["message"])
		})
	}
	valid, err := jwtService.GenerateToken(userId, auth.TokenAccess, auth.ExpiresInOneDay)
	s.Require().NoError(err)

	s.Run("success", func() {
		w := s.newRequest(s.ctx, "GET", url, nil, valid)

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
