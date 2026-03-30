package ratelimit

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter interface {
	Allow(key string) Result
}

type Result struct {
	Allowed   bool
	Remaining int
	ResetAt   time.Time
	RetryAt   time.Time
}

type TokenBucket struct {
	sync.Mutex
	rate       float64 // tokens per second
	buckets    map[string]*bucket
	bucketSize int
}

type bucket struct {
	tokens   float64
	lastFill time.Time
}

func NewTokenBucket(ratePerSecond float64, bucketSize int) *TokenBucket {
	return &TokenBucket{
		rate:       ratePerSecond,
		bucketSize: bucketSize,
		buckets:    make(map[string]*bucket),
	}
}

func (tb *TokenBucket) Allow(key string) Result {
	tb.Lock()
	defer tb.Unlock()

	now := time.Now()
	b, exists := tb.buckets[key]
	// Create new bucket if there is none
	if !exists {
		b = &bucket{tokens: float64(tb.bucketSize), lastFill: now}
		tb.buckets[key] = b
	}
	// Fill bucket
	elapsedSeconds := now.Sub(b.lastFill).Seconds()
	b.tokens = elapsedSeconds * tb.rate
	if b.tokens > float64(tb.bucketSize) {
		b.tokens = float64(tb.bucketSize)
	}
	b.lastFill = now

	// There is no tokens in the bucket
	if b.tokens < 1 {
		waitSeconds := (1 - b.tokens) / tb.rate
		return Result{
			Allowed:   false,
			Remaining: 0,
			RetryAt:   now.Add(time.Duration(waitSeconds * float64(time.Second))),
		}
	}
	b.tokens--

	return Result{
		Allowed:   true,
		Remaining: int(b.tokens),
	}
}
func KeyFunc(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	return r.RemoteAddr
}
