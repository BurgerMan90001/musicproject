package server

import (
	"fmt"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/middleware"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/service/auth"
)

type Server struct {
	mux  *http.ServeMux
	repo repository.Repository
	cfg  config.Config
}

func New(mux *http.ServeMux, repo repository.Repository, cfg config.Config) *Server {
	return &Server{mux, repo, cfg}
}

func (h *Server) Handle() {
	authService := auth.New(h.cfg.API.Auth, h.repo)

	// setup routes
	h.handleFunc("/health", handleHealth)
	// use routes
	h.handleFunc("/user/{id}", handler.HandleUserID(h.repo))

	h.handleFunc("/songs/{id}", handler.HandleSongs(h.repo))
	h.handleFunc("/artists/{id}", handler.HandleArtists())

	// auth routes
	h.handleFunc("/auth/login", handler.HandleLogin(authService, h.repo))
	h.handleFunc("/auth/signup", handler.HandleSignup(authService, h.repo))
	h.handleFunc("/auth/refresh", handler.HandleRefresh())
	h.handleFunc("/auth/email/reset", handler.HandleEmailReset())

	h.handleFunc("/auth/google/login", handler.HandleOauthLogin(authService))
	h.handleFunc("/auth/google/redirect", handler.HandleOauthGoogleRedirect(authService))

	h.handleFunc("/secret", middleware.JWTMiddleware(nil, handleSecret))
	// static file server
	//mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//middleware.Logger(mux)
}
func (s *Server) handleFunc(route string, function http.HandlerFunc) {
	s.mux.HandleFunc(s.URL(route), function)
}
func (s *Server) URL(route string) string {
	return fmt.Sprintf("/%s/%s", s.cfg.API.Version, route)
}

/*
	func (h *Handler) Cleanup() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		// go func() {
		// 	<-c
		// 	// stop db
		// 	if err := h.repo.Stop(); err != nil {
		// 		panic(err)
		// 	}
		// 	os.Exit(1)
		// }()
	}
*/
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "alive")
}

func handleSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
