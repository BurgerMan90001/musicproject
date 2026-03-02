package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"okapi.com/internal/auth"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
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

func JWTMiddleware(jwtKey []byte, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := auth.ParseToken(jwtKey, r)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid token:", err)
			return
		}

		fmt.Fprintln(w, "Valid token")
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
