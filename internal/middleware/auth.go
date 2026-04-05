package middleware

import (
	"context"
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
)

func WithAuth(jwtService *auth.JWTService, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(auth.TokenRefresh)
		if err != nil {
			jsonutil.WriteError(w, auth.ErrNoAccessToken, http.StatusUnauthorized)
			return
		}

		claims, err := jwtService.ParseAccessToken(cookie.Value)
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusUnauthorized)
			return
		}

		// Pass claims to the next handler
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
