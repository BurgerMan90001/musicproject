package middleware

import (
	"context"
	"errors"
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

type validator interface {
	Validate(ctx context.Context, tokenString string) (*model.Claims, error)
}

func RequireAuth(validator validator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(auth.TokenAccess)
			if err != nil {
				jsonutil.WriteError(w, auth.ErrNoAccessToken, http.StatusUnauthorized)
				return
			}
			ctx := r.Context()

			claims, err := validator.Validate(ctx, cookie.Name)
			if err != nil {
				jsonutil.WriteError(w, errors.New("Unauthorized"), http.StatusUnauthorized)
				return
			}

			// Pass claims to the next handler
			ctxClaims := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctxClaims))
		})
	}
}
