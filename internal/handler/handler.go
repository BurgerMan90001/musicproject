package handler

import (
	"context"
	"database/sql"
	"net/http"

	"musicproject.com/internal/config"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/middleware"
	"musicproject.com/internal/middleware/ratelimit"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/services/auth"
	"musicproject.com/internal/services/file"
	"musicproject.com/internal/services/upload"
	"musicproject.com/pkg/model"
)

func NewMux(ctx context.Context, cfg *config.Config, store file.Blobstore, db *sql.DB) (http.Handler, error) {
	mux := http.NewServeMux()

	userRepo := postgres.NewUser(db)
	songRepo := postgres.NewSong(db)
	refreshTokenRepo := postgres.NewRefreshToken(db)
	//ratingRepo := postgres.NewRating(db)

	authService, err := auth.New(ctx, cfg.Auth, refreshTokenRepo, userRepo)
	if err != nil {
		return nil, err
	}
	uploadService := upload.New(cfg.Upload.Bucket, false, store, songRepo)

	// emailService, err := email.New()
	// if err != nil {
	// 	return nil, err
	// }

	mux.HandleFunc("/", handleNotFound())
	mux.HandleFunc("/health", handleHealth)

	// auth routes
	mux.HandleFunc("/auth/login", handleLogin(authService))
	mux.HandleFunc("/auth/signup", handleSignup(authService))
	mux.HandleFunc("/auth/refresh", handleRefresh(authService))
	mux.HandleFunc("/auth/logout", handleLogout(authService))
	//mux.HandleFunc("/auth/reset", authHandler.handleEmailReset(emailService))

	// oauth routes
	mux.HandleFunc("/auth/google/login", handleOauthLogin(authService.Google))
	mux.HandleFunc("/auth/google/redirect", handleOauthRedirect(authService.Google))

	// Metadata routes
	mux.HandleFunc("/users", handleUsers(userRepo))
	mux.HandleFunc("/users/{id}", handleUsersID(userRepo))

	mux.HandleFunc("/songs/{id}", handleGetSongsMetadata(songRepo))

	// File routes
	mux.HandleFunc("/upload/songs", handleSongUpload(uploadService))
	//mux.HandleFunc("/files/audio/{id}", handleAudio(songService))

	// MAYBE
	//mux.HandleFunc("/songs/{id}/rating", handleSongRating())

	// MAYBE
	//mux.HandleFunc("/artists/{id}", HandleArtists())

	// Uploads audio encoded
	//mux.HandleFunc("POST /files/audio/encode", HandleAudioEncode())

	// Image
	mux.HandleFunc("/files/image", handleImage())

	// Test routes
	mux.Handle("/protected", middleware.RequireAuth(authService)(handleTest()))

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
	root.Handle("/", handleNotFound())
	return root, nil
}

func handleNotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonutil.WriteError(w, &model.Error{
			Code:    http.StatusNotFound,
			Message: "route not found",
		})
	})

}

func handleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		ctx := r.Context()
		claims, ok := contextClaims(ctx)
		if ok {
			jsonutil.WriteJSON(w, claims, http.StatusOK)
		}
		jsonutil.WriteJSON(w, nil, http.StatusOK)
	}
}
