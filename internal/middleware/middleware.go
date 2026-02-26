package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"movieexample.com/internal/auth"
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
func JWTMiddleware(jwtKey []byte, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		}, request.WithClaims(&auth.Claims{}))

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprintln(w, "Invalid token:", err)
			return
		}

		fmt.Fprintln(w, "Valid token")
		next.ServeHTTP(w, r)
	})
}
