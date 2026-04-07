package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/middleware"
	"musicproject.com/internal/middleware/ratelimit"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/services/auth"
	"musicproject.com/internal/services/email"
	"musicproject.com/internal/services/file"
	"musicproject.com/internal/services/song"
)

func NewMux(ctx context.Context, cfg *config.Config, db *sql.DB, sm secrets.Manager) (http.Handler, error) {
	mux := http.NewServeMux()

	userRepo := postgres.NewUser(db)
	songRepo := postgres.NewSong(db)

	//ratingRepo := postgres.NewRating(db)
	// store, err := file.NewS3(ctx)
	store := file.NewFileSystem()
	// if err != nil {
	// 	return nil, err
	// }

	authService, err := auth.New(ctx, cfg.Services.Auth, userRepo, sm)
	if err != nil {
		return nil, err
	}
	userHandler := &userHandler{userRepo: userRepo}
	authHandler := &authHandler{authService: authService}

	songService := song.NewSong(store, songRepo)
	//store := file.NewFileSystem()

	emailService, err := email.New(ctx, sm)
	if err != nil {
		return nil, err
	}

	// setup routes
	mux.HandleFunc("/", HandleNotFound)
	mux.HandleFunc("/health", HandleHealth)

	mux.HandleFunc("/users", userHandler.handleUsers())
	mux.HandleFunc("/users/{id}", userHandler.handleUsersID())

	mux.HandleFunc("/songs/{id}", HandleSongsMetadata(songRepo))
	mux.Handle("POST /songs/upload", HandleSongUpload(songService))

	// maybe
	//mux.HandleFunc("/artists/{id}", HandleArtists())

	// auth routes
	mux.HandleFunc("/auth/login", authHandler.handleLogin())
	mux.HandleFunc("/auth/signup", authHandler.handleSignup())

	mux.HandleFunc("/auth/refresh", authHandler.handleRefresh())
	mux.HandleFunc("/auth/reset", authHandler.handleEmailReset(emailService))

	// oauth routes
	mux.HandleFunc("/auth/google/login", HandleOauthLogin(authService.Google))
	mux.HandleFunc("/auth/google/redirect", HandleOauthGoogleRedirect(authService.Google))

	// Test routes
	mux.Handle("/protected", middleware.RequireAuth(authService.JWT, HandleTest))
	mux.HandleFunc("/audio", handleAudio(songService))

	// file server

	var handler http.Handler = mux
	if cfg.Middleware.Logger {
		handler = middleware.Logger(handler)
	}
	if cfg.Middleware.Ratelimit {
		rl := ratelimit.NewTokenBucket(15, 30)
		handler = middleware.RateLimit(rl, handler)
	}
	root := http.NewServeMux()

	root.Handle("/v1/", http.StripPrefix("/v1", handler))
	//root.Handle("/", http.HandlerFunc(HandleNotFound))

	// var handler http.Handler = root
	// if cfg.Middleware.Logger {
	// 	handler = middleware.WithLogger(root)
	// }
	return root, nil
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	jsonutil.WriteError(w, errors.New("route not found"), http.StatusNotFound)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ctx := r.Context()
	claims, ok := contextClaims(ctx)
	if ok {
		jsonutil.WriteJSON(w, claims, http.StatusOK)
	}
	jsonutil.WriteJSON(w, nil, http.StatusOK)
}
