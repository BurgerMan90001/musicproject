package handler

import (
	"fmt"
	"log"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/internal/services/file"
)

type Server struct {
	mux *http.ServeMux
	cfg *config.Config
}

func NewServer(cfg *config.Config, repo repository.Repository) *Server {
	mux := http.NewServeMux()

	authService := auth.New(cfg.Services.Auth, repo)
	fileService := file.New()

	//mux.Handle("")
	// setup routes
	mux.HandleFunc("/v1/health", HandleHealth)
	// use routes
	mux.HandleFunc("/v1/user/{id}", HandleUserID(repo))

	mux.HandleFunc("/v1/songs/{id}", HandleSongs(repo))
	mux.HandleFunc("/v1/songs/upload", HandleSongUpload(fileService))
	mux.HandleFunc("/v1/artists/{id}", HandleArtists())

	// auth routes
	mux.HandleFunc("/v1/auth/login", HandleLogin(authService))
	mux.HandleFunc("/v1/auth/signup", HandleSignup(authService))
	mux.HandleFunc("/v1/auth/refresh", HandleRefresh(authService))
	mux.HandleFunc("/v1/auth/reset", HandleEmailReset())

	// oauth routes
	mux.HandleFunc("/v1/auth/google/login", HandleOauthLogin(authService.Google))
	mux.HandleFunc("/v1/auth/google/redirect", HandleOauthGoogleRedirect(authService))

	mux.HandleFunc("/v1/protected", AuthMiddleware(authService, HandleTest))
	// static file server
	mux.Handle("/static", http.FileServer(http.Dir("public")))

	//root := http.NewServeMux()

	//root.Handle("/v1/", http.StripPrefix("/v1/", mux))

	s := &Server{mux, cfg}

	// add middleware
	//middleware.Logger(mux)
	return s
}
func (s *Server) root() string {
	return fmt.Sprintf("/v%s", s.cfg.API.Version)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

//	func (s *Server) URL(route string) string {
//		return fmt.Sprintf("/%s%s", s.cfg.API.Version, route)
//	}
func (s *Server) Listen() {
	log.Printf("Server listening at %s", s.cfg.URL())
	if err := http.ListenAndServe(s.cfg.URL(), s.mux); err != nil {
		panic(err)
	}
}

// func (s *Server) handleFunc(route string, function http.HandlerFunc) {
// 	s.mux.HandleFunc(s.URL(route), function)
// }

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "alive")
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	WriteJSON(w, nil, http.StatusOK)
}

/*
func Cleanup() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// stop db
		if err := h.repo.Stop(); err != nil {
			panic(err)
		}
		os.Exit(1)
	}()
}
*/
