package handler

import (
	"net/http"

	"musicproject.com/internal/services/auth"
)

func accessCookie(value string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     auth.TokenRefresh,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/v1/auth/refresh",
		MaxAge:   maxAge,
	}
}

func refreshCookie(value string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     auth.TokenRefresh,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/v1/auth/refresh",
		MaxAge:   maxAge,
	}
}
