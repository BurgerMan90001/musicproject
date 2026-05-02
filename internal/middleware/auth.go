package middleware

import (
	"context"
	"net/http"

	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/auth"
	"songsled.com/pkg/model"
)

type validateFunc func(ctx context.Context, tokenString string, needRoles ...string) (*model.Claims, error)

type Auth struct {
	validate validateFunc
}

func NewAuth(validate validateFunc) *Auth {
	return &Auth{validate: validate}
}
func (m *Auth) RequireAuth(needRoles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(string(model.TokenAccess))
			if err != nil {
				jsonutil.WriteError(w, auth.ErrInvalidToken(err))
				return
			}

			claims, err := m.validate(r.Context(), cookie.Value, needRoles...)
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
