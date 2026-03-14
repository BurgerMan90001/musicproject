package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

var jwtSecret = []byte("test")

func TestJWTMiddleware(t *testing.T) {
	validToken, err := auth.GenerateToken(jwtSecret, &model.User{
		ID: uuid.Nil,
	}, auth.TokenAccess, auth.ExpiresInOneDay)
	if err != nil {
		t.Error(err)
	}
	expiredToken, err := auth.GenerateToken(jwtSecret, &model.User{
		ID: uuid.Nil,
	}, auth.TokenAccess, time.Now().Add(time.Hour*-1))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name       string
		token      string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "valid token",
			token:      validToken,
			wantStatus: http.StatusOK,
		},
		{
			name:       "expired token",
			token:      expiredToken,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "invalid token",
			token:      "eyJhbGcaaaaI1NiIsInR5cCI6IkpXVCJ9.e30.P4Lqll22jQQJ1eMJikvNg5HKG-cKB0hUZA9BZFIG7Jk",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "empty token or no header",
			token:      "",
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			w := httptest.NewRecorder()

			handler := JWTMiddleware(jwtSecret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(w, req)

			// body, err := io.ReadAll(w.Body)
			// if err != nil {
			// 	t.Error(err)
			// }
			assert.Equal(t, tt.wantStatus, w.Code, tt.name)
			//assert.NotEmpty()
			//assert.Equal(t, tt.wantBody, string(body), tt.name)
		})
	}
}
