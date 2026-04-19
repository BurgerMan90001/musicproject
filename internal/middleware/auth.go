package middleware

import (
	"context"
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

type authenticator interface {
	Validate(tokenString string) (*model.Claims, error)
}

func RequireAuth(validator authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(string(model.TokenAccess))
			if err != nil {
				jsonutil.WriteError(w, auth.ErrNoAccessToken)
				return
			}

			claims, err := validator.Validate(cookie.Value)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			// Pass claims to the next handler
			ctxClaims := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctxClaims))
		})
	}
}
