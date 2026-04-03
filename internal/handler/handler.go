package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/handler/middleware/ratelimit"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/services/auth"
	"musicproject.com/internal/services/file"
)

func NewMux(ctx context.Context, cfg *config.Config, db *sql.DB) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	// setup middleware
	var rl ratelimit.RateLimiter
	if cfg.Middleware.Ratelimit {
		rl = ratelimit.NewTokenBucket(15, 30)
	}

	userRepo := postgres.NewUser(db)
	songRepo := postgres.NewSong(db)
	api := Logger(PanicRecovery(RateLimitMiddleware(rl, mux)))
	//ratingRepo := postgres.NewRating(db)

	store, err := file.NewS3(ctx)
	if err != nil {
		return nil, err
	}

	authService := auth.New(cfg.Services.Auth, userRepo)
	fileService := file.NewSong(store)

	// setup routes
	mux.HandleFunc("/", HandleNotFound)
	mux.HandleFunc("/health", HandleHealth)

	mux.HandleFunc("/users", HandleUsers(userRepo))
	mux.HandleFunc("/users/{id}", HandleUsersID(userRepo))

	mux.HandleFunc("/songs/{id}", HandleSongs(songRepo))
	mux.HandleFunc("POST /songs/upload", HandleSongUpload(fileService))

	mux.HandleFunc("/artists/{id}", HandleArtists())

	// auth routes
	mux.HandleFunc("/auth/login", HandleLogin(authService))
	mux.HandleFunc("/auth/signup", handleSignup(authService))

	mux.HandleFunc("/auth/refresh", HandleRefresh(authService))
	mux.HandleFunc("/auth/reset", HandleEmailReset())

	// oauth routes
	mux.HandleFunc("/auth/google/login", HandleOauthLogin(authService.Google))
	mux.HandleFunc("/auth/google/redirect", HandleOauthGoogleRedirect(authService.Google))

	mux.HandleFunc("/protected", AuthMiddleware(authService.JWT, HandleTest))
	// static file server
	mux.Handle("/static", http.FileServer(http.Dir("public")))

	//var api http.Handler = mux

	root := http.NewServeMux()

	root.Handle("/v1/", http.StripPrefix("/v1", api))
	root.Handle("/", http.HandlerFunc(HandleNotFound))
	//test := root.(*http.Handler)
	return root, nil
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	WriteError(w, errors.New("route not found"), http.StatusNotFound)
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
