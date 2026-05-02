package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/redis/go-redis/v9"
	"songsled.com/internal/config"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/middleware"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/auth"
	"songsled.com/internal/services/file"
	"songsled.com/internal/services/search"
	"songsled.com/internal/services/upload"
	"songsled.com/pkg/model"
)

func New(
	ctx context.Context,
	cfg *config.Config,
	store file.Blobstore,
	repo *postgres.Repo,
	rdb *redis.Client,
) (http.Handler, error) {
	root := chi.NewRouter()

	userRepo := postgres.NewUser(repo.Queries)
	songRepo := postgres.NewSong(repo.Queries)
	playlistRepo := postgres.NewPlaylist(repo.Queries)

	searchService := search.NewPostgres()
	authService, err := auth.New(ctx, cfg.Auth, rdb, userRepo)
	if err != nil {
		return nil, err
	}

	uploadService, err := upload.NewSong(cfg.Upload.Bucket, "audio",
		false, true, 30*time.Minute, store, songRepo)
	if err != nil {
		return nil, err
	}

	authMw := middleware.NewAuth(authService.Validate)
	// rl := ratelimit.NewTokenBucket(15, 30)

	if !testing.Testing() {
		root.Use(middleware.Logger())
		root.Use(middleware.Limit(nil))

	}
	root.Use(cors.Handler(cors.Options{

		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	root.Route("/v1", func(api chi.Router) {

		api.Route("/auth", func(r chi.Router) {
			r.HandleFunc("/signup", handleSignup(authService))
			r.HandleFunc("/login", handleLogin(authService))
			r.HandleFunc("/refresh", handleRefresh(authService))
			r.HandleFunc("/logout", handleLogout(authService))
			//mux.HandleFunc("/auth/reset", authHandler.handleEmailReset(emailService))

			// oauth routes
			r.HandleFunc("/google/login", handleOauthLogin(authService.Google))
			r.HandleFunc("/google/redirect", handleOauthRedirect(authService.Google))
		})

		api.Route("/users", func(r chi.Router) {
			r.Get("/", handleUsers(userRepo))

			r.Get("/{id}", handleGetUsersId(userRepo))
			r.Delete("/{id}", handleDelteUsersId(userRepo))

			r.Get("/{id}/history", handleUserHistory())
		})

		api.Route("/songs", func(r chi.Router) {
			r.HandleFunc("/", handleSongs(searchService, songRepo))
			r.HandleFunc("/{id}", handleGetSong(songRepo))

			r.HandleFunc("/upload", handleSongUpload(uploadService))

		})
		api.Route("/playlists", func(r chi.Router) {
			r.HandleFunc("/", handlePlaylists(playlistRepo))
			r.HandleFunc("/{id}", handlePlaylistsID(playlistRepo))
		})

		api.Route("/admin", func(r chi.Router) {
			r.Use(authMw.RequireAuth(auth.RoleAdmin))
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				jsonutil.WriteJSON(w, nil, http.StatusOK)
			})

		})
	})

	root.Get("/health", handleHealth)

	root.NotFound(func(w http.ResponseWriter, r *http.Request) {
		jsonutil.WriteError(w, &model.Error{
			Code:    http.StatusNotFound,
			Message: "Route not found",
		})
	})
	root.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		jsonutil.WriteError(w, &model.Error{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		})
	})

	return root, nil
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
