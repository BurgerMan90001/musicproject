package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"musicproject.com/internal/handler/ratelimit"
	"musicproject.com/internal/services"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, req)
	})
}

func AuthMiddleware(authService services.JWT, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(AccessCookie)
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				WriteError(w, auth.ErrNoAccessToken, http.StatusUnauthorized)
			default:
				WriteError(w, err, http.StatusBadRequest)
			}
			return
		}

		claims, err := authService.ParseAccessToken(cookie.Value)
		if err != nil {
			WriteError(w, err, http.StatusUnauthorized)
			return
		}

		if claims.TokenType != auth.TokenAccess {
			WriteError(w, auth.ErrInvalidTokenType, http.StatusUnauthorized)
			return
		}

		// Pass claims to the next handler
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func RateLimitMiddleware(rl ratelimit.RateLimiter, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := ratelimit.KeyFunc(r)
		result := rl.Allow(key)

		w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(result.Remaining))
		if !result.ResetAt.IsZero() {
			w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(result.ResetAt.Unix(), 10))
		}
		if !result.Allowed {
			w.Header().Set("Retry-After", fmt.Sprintf("%.0f", time.Until(result.RetryAt).Seconds()))
			WriteError(w, ErrRateLimit, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
func chainMiddleware(mw func(http.Handler) http.Handler) {

}
func contextClaims(ctx context.Context) (*model.Claims, bool) {
	claims, ok := ctx.Value("claims").(*model.Claims)
	return claims, ok
}
