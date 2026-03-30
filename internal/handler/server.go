package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/handler/ratelimit"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/internal/services/file"
	"musicproject.com/pkg/model"
)

type Server struct {
	*http.ServeMux
	cfg *config.Config
}

func NewServer(cfg *config.Config, repo repository.Repository) *Server {
	mux := http.NewServeMux()

	rl := ratelimit.NewTokenBucket(15, 30)

	authService := auth.New(cfg.Services.Auth, repo)
	fileService := file.New()

	// setup routes
	mux.HandleFunc("/", HandleNotFound)
	mux.HandleFunc("/health", HandleHealth)

	http.HandleFunc("/users", HandleUsers(repo))
	mux.HandleFunc("/users/{id}", HandleUsersID(repo))

	mux.HandleFunc("/songs/{id}", HandleSongs(repo))
	mux.HandleFunc("/songs/upload", HandleSongUpload(fileService))

	mux.HandleFunc("/artists/{id}", HandleArtists())

	// auth routes
	mux.HandleFunc("/auth/login", HandleLogin(authService))
	mux.HandleFunc("/auth/signup", HandleSignup(authService))

	mux.HandleFunc("/auth/refresh", HandleRefresh(authService.JWT))
	mux.HandleFunc("/auth/reset", HandleEmailReset())

	// oauth routes
	mux.HandleFunc("/auth/google/login", HandleOauthLogin(authService.Google))
	mux.HandleFunc("/auth/google/redirect", HandleOauthGoogleRedirect(authService.Google))

	mux.HandleFunc("/protected", AuthMiddleware(authService.JWT, HandleTest))
	// static file server
	mux.Handle("/static", http.FileServer(http.Dir("public")))

	var api http.Handler = mux
	if cfg.Middleware.Ratelimit {
		api = Logger(PanicRecovery(RateLimitMiddleware(rl, mux)))
	}

	root := http.NewServeMux()

	root.Handle("/v1/", http.StripPrefix("/v1", api))
	root.Handle("/", http.HandlerFunc(HandleNotFound))

	s := &Server{root, cfg}

	return s
}

func (s *Server) url() string {
	return fmt.Sprintf("%v:%d", s.cfg.API.Host, s.cfg.API.Port)
}

func (s *Server) Run() {
	
}
func (s *Server) Listen() {
	log.Printf("Server listening at %s. v1", s.url())
	if err := http.ListenAndServe(s.url(), s); err != nil {
		panic(err)
	}
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, model.Health{
		Message: "alive",
	}, http.StatusOK)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	WriteError(w, errors.New("not found"), http.StatusNotFound)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ctx := r.Context()
	claims, ok := contextClaims(ctx)
	if ok {
		WriteJSON(w, claims, http.StatusOK)
	}
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
