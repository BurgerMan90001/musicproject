package handler

import (
	"net/http"
	"testing"
)

func TestAuthHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		wantStatus int
	}{
		{
			name:       "successful login",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "successful signup",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Skip("not implemented")
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
