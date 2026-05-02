package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"songsled.com/internal/jsonutil"
	"songsled.com/internal/middleware/ratelimit"
)

type RateLimit struct {
	rl ratelimit.RateLimiter
}

func NewRateLimit(rl ratelimit.RateLimiter) *RateLimit {
	return &RateLimit{rl}
}

func Limit(rl ratelimit.RateLimiter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rl == nil {
				next.ServeHTTP(w, r)
				return
			}
			key := ratelimit.KeyFunc(r)

			result := rl.Allow(key)

			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(result.Remaining))
			if !result.ResetAt.IsZero() {
				w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(result.ResetAt.Unix(), 10))
			}
			if !result.Allowed {
				w.Header().Set("Retry-After", fmt.Sprintf("%.0f", time.Until(result.RetryAt).Seconds()))
				jsonutil.WriteError(w, ratelimit.ErrRateLimit)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
