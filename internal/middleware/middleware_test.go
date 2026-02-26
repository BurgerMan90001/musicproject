package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTMiddleware(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		token      string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "valid token",
			token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.P4Lqll22jQQJ1eMJikvNg5HKG-cKB0hUZA9BZFIG7Jk",
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid token",
			token:      "paaaah123",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "empty token",
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

			handler := JWTMiddleware([]byte("test"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("wrong status got: %v wanted: %v", w.Code, tt.wantStatus)
			}

		})
	}
}
