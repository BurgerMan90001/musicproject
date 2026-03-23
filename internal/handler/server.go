package handler

import (
	"fmt"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/service/auth"
)

type Server struct {
	mux *http.ServeMux
	cfg config.Config
}

func New(mux *http.ServeMux, cfg config.Config) *Server {

	s := &Server{mux, cfg}

	repo := postgres.New(s.cfg)
	authService := auth.New(s.cfg.Services.Auth, repo.User)

	// setup routes
	s.handleFunc("/health", HandleHealth)
	// use routes
	s.handleFunc("/user/{id}", HandleUserID(repo.User))

	s.handleFunc("/songs/{id}", HandleSongs(repo.Song))
	s.handleFunc("/artists/{id}", HandleArtists())

	// auth routes
	s.handleFunc("/auth/login", HandleLogin(authService))
	s.handleFunc("/auth/signup", HandleSignup(authService))
	s.handleFunc("/auth/refresh", HandleRefresh())
	s.handleFunc("/auth/email/reset", HandleEmailReset())

	s.handleFunc("/auth/google/login", HandleOauthLogin(authService))
	s.handleFunc("/auth/google/redirect", HandleOauthGoogleRedirect(authService))

	s.handleFunc("/secret", AuthMiddleware(authService, HandleTest))
	// static file server
	//mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//middleware.Logger(mux)
	return s
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
func (s *Server) URL(route string) string {
	return fmt.Sprintf("/%s/%s", s.cfg.API.Version, route)
}

func (s *Server) handleFunc(route string, function http.HandlerFunc) {
	s.mux.HandleFunc(s.URL(route), function)

}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "alive")
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
