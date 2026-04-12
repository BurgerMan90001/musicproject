package model

import "net/http"

type TokenType string

const (
	TokenAccess  TokenType = "accessKey"
	TokenRefresh TokenType = "refreshKey"
)

// Max age is in seconds
func (t TokenType) Cookie(value string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     string(t),
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/v1/auth/refresh",
		MaxAge:   maxAge,
	}
}

func (t TokenType) Clear() *http.Cookie {
	return t.Cookie("", -1)
}
