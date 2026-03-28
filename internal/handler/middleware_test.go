package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func TestAuthMiddleware(t *testing.T) {
	url := "/v1/protected"
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
		t.Run(tt.name, testCase(func(t *testing.T, c *testContext) {
			tokenString, err := c.authService.GenerateToken(tt.userId, tt.tokenType, tt.tokenExpireAt)
			assert.NoError(t, err)

			w, err := newRequest(ctx, "GET", url, nil, tokenString)
			assert.NoError(t, err)

			resBody, err := model.ReadJSON[model.Response](w.Result().Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantMessage, resBody.Message)
		}))
	}

	t.Run("invalid token", testCase(func(t *testing.T, c *testContext) {

		w, err := newRequest(ctx, "GET", url, nil, "aimdlwdnaljngrlgns")
		assert.NoError(t, err)

		resBody, err := model.ReadJSON[model.Response](w.Result().Body)
		assert.NoError(t, err)

		assert.Equal(t, jwt.ErrTokenMalformed.Error(), resBody.Message)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}))
	t.Run("empty token or no header", testCase(func(t *testing.T, c *testContext) {
		w, err := newRequest(ctx, "GET", url, nil, "")
		assert.NoError(t, err)

		resBody, err := model.ReadJSON[model.Response](w.Result().Body)
		assert.NoError(t, err)

		assert.Equal(t, auth.ErrNoAccessToken.Error(), resBody.Message)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}))
}
