package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"musicproject.com/config"
	"musicproject.com/internal/service/auth"
)

var jwtSecret = []byte("test")

func TestJWTMiddleware(t *testing.T) {
	cfg := config.ReadConfigFile()
	authService := auth.New(cfg)

	validToken, err := authService.GenerateToken(uuid.Nil, auth.TokenAccess, auth.ExpiresInOneDay)
	if err != nil {
		t.Error(err)
	}
	expiredToken, err := authService.GenerateToken(uuid.Nil,
		auth.TokenAccess, time.Now().Add(time.Hour*-1))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name     string
		token    string
		wantRes  string
		wantCode int
		//wantErr  map[string]error
	}{

		{
			name:     "valid token",
			wantCode: http.StatusOK,
			token:    validToken,
		},
		{
			name:  "expired token",
			token: expiredToken,

			wantRes:  jwt.ErrTokenExpired.Error(),
			wantCode: http.StatusUnauthorized,
		},
		{

			token:    "eyJhbGcaaaaaaaaaaaaaaaaaaaaaaaJk",
			name:     "invalid token",
			wantCode: http.StatusUnauthorized,
		},
		{
			token:    "",
			name:     "empty token or no header",
			wantCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/protected", nil)
			if tt.token != "" {
				r.Header.Set("Authorization", "Bearer "+tt.token)
			}
			w := httptest.NewRecorder()

			handler := JWTMiddleware(jwtSecret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(w, r)

			// res, err := ReadJSON[map[string]string](w.Result().Body)
			// if err != nil {
			// 	t.Error(err)
			// }

			assert.Equal(t, tt.wantCode, w.Code, tt.name)
			//assert.JSONEqf()

			//assert.Equal(t, tt.wantRes, res["data"], tt.name)
			//assert.Equal(t, tt.wantRes, res["error"], tt.name)
		})
	}
}
