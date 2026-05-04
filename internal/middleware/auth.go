package middleware

import (
	"context"
	"net/http"

	"songsled.com/pkg/model"
)

// type validateFunc func(ctx context.Context, tokenString string, needRoles ...string) (*model.Claims, error)

//	type Auth struct {
//		validate validateFunc
//	}
type validator interface {
	Validate(ctx context.Context, token string) (*model.Claims, error)
	// RedirectUrl() string
}

func RequireAuth(v validator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// cookie, err := r.Cookie(string(model.TokenAccess))
			// if err != nil {
			// 	jsonutil.WriteError(w, auth.ErrInvalidToken(err))
			// 	return
			// }

			
			// claims, err := v.Validate(r.Context(), cookie.Value)
			// if err != nil {

			// 	jsonutil.WriteError(w, err)

			// 	return
			// }

			// // Pass claims to the next handler
			// ctxClaims := context.WithValue(r.Context(), "claims", claims)
			// next.ServeHTTP(w, r.WithContext(ctxClaims))
		})
	}
}
