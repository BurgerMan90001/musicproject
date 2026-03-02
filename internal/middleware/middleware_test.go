package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var jwtSecret = "test"

func TestJWTMiddleware(t *testing.T) {
	tests := []struct {
		name  string
		token string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "valid token",
			token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.P4Lqll22jQQJ1eMJikvNg5HKG-cKB0hUZA9BZFIG7Jk",
			wantStatus: http.StatusOK,
		},
		{
			name:       "valid oauth token",
			token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJva2FwaSIsImV4cCI6MTc3MjM5OTM4NX0.P0f93OMwTD180Tr9rEctXOtqeBI4Zi83wwLbXy85gOc",
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid token",
			token:      "eyJhbGcaaaaI1NiIsInR5cCI6IkpXVCJ9.e30.P4Lqll22jQQJ1eMJikvNg5HKG-cKB0hUZA9BZFIG7Jk",
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

			handler := JWTMiddleware([]byte(jwtSecret), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("wrong status got: %v wanted: %v", w.Code, tt.wantStatus)
			}
		})
	}
}
