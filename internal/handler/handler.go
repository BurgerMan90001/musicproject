package handler

import (
	"fmt"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/controller/user"
	"musicproject.com/internal/middleware"
	"musicproject.com/internal/repository"
)

type Handler struct {
	mux  *http.ServeMux
	repo repository.Repository
	cfg  config.Config
}

func New(mux *http.ServeMux, cfg config.Config) *Handler {
	repo := cfg.NewRepository()
	return &Handler{mux, repo, cfg}
}

func (h *Handler) Register(path string) {
	userController := user.New(h.repo)
	jwtKey := h.cfg.JWTAccessKey()
	oathCfg := h.cfg.GoogleOathConfig()

	// setup routes
	h.mux.HandleFunc("/health", handleHealth)
	// user routes
	h.mux.HandleFunc("/user", handleUser(userController))

	h.mux.HandleFunc("/songs", handleSongs())

	h.mux.HandleFunc("/artists", handleArtists())

	// auth routes
	h.mux.HandleFunc("/auth/login", handleLogin(jwtKey, userController))
	h.mux.HandleFunc("/auth/signup", handleSignup(jwtKey, userController))
	h.mux.HandleFunc("/auth/refresh", handleRefresh())

	h.mux.HandleFunc("/auth/google/login", handleOauthGoogleLogin(oathCfg))
	h.mux.HandleFunc("/auth/google/redirect", handleOauthGoogleRedirect(jwtKey, oathCfg))

	h.mux.HandleFunc("/secret", middleware.JWTMiddleware(jwtKey, handleSecret))
	// static file server
	//mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//middleware.Logger(mux)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "alive")
}

func handleSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleSongs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

		case http.MethodPut:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}
