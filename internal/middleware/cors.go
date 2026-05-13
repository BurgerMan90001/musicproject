package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func Cors() func(http.Handler) http.Handler {

	var orgins = []string{"http://localhost:5173", "https://songsled.com"}
	// if os.Getenv("ENV") == "prod" {
	// 	orgins = []string{"https://songsled.com"}
	// }
	return cors.Handler(cors.Options{
		AllowedOrigins:   orgins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
