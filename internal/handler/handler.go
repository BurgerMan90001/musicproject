package handler

import (
	"fmt"
	"net/http"

	"okapi.com/config"
	"okapi.com/internal/controller/user"
	"okapi.com/internal/middleware"
	"okapi.com/internal/repository"
)

type Handler struct {
	mux  *http.ServeMux
	repo repository.Repository
	cfg  config.Config
}

func New(mux *http.ServeMux, cfg config.Config) *Handler {
	repo := cfg.NewRepository()
	return &Handler{mux,repo,cfg,}
}

func (h *Handler) Register(path string) {
	userController := user.New(h.repo)
	jwtKey := h.cfg.JWTAccessKey()
	oathCfg := h.cfg.GoogleOathConfig()
	
	// setup routes
	h.mux.HandleFunc("/health", handleHealth)
	// user routes
	h.mux.HandleFunc("/user", handleUser(userController))

	// auth routes
	h.mux.HandleFunc("/auth/login", handleLogin(jwtKey, userController))
	h.mux.HandleFunc("/auth/signup", handleSignup(jwtKey, userController))

	h.mux.HandleFunc("/auth/google/login", handleOathGoogleLogin(oathCfg))
	h.mux.HandleFunc("/auth/google/redirect", handleOathGoogleRedirect(jwtKey, oathCfg))

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
