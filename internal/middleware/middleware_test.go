package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTMiddleware(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		method     string
		url        string
		body       io.Reader
		token      string
		wantBody   string
		wantStatus int
	}{
		{
			name:       "success",
			method:     "GET",
			url:        "localhost:8081",
			token:      "success",
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, tt.body)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			rr := httptest.NewRecorder()

			handler := JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			if rr.Code != tt.wantStatus {
				t.Errorf("wrong status got: %v wanted: %v", rr.Code, tt.wantStatus)
			}

			handler.ServeHTTP(rr, req)
		})
	}
}
