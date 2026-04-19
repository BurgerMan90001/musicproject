package ratelimit

import (
	"net/http"
	"time"

	"musicproject.com/pkg/model"
)

var ErrRateLimit = &model.Error{
	Code:    http.StatusTooManyRequests,
	Message: "Rate limit exceeded",
}

type RateLimiter interface {
	Allow(key string) Result
}

type Result struct {
	Allowed   bool
	Remaining int
	ResetAt   time.Time
	RetryAt   time.Time
}

func KeyFunc(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	return r.RemoteAddr
}
