package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"songsled.com/internal/config"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/middleware"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/file"
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

	artistRepo := postgres.NewArtistRepo(repo.Queries)
	songRepo := postgres.NewSongRepo(repo.Queries)
	albumRepo := postgres.NewAlbumRepo(repo.Queries)
	playlistRepo := postgres.NewPlaylistRepo(repo.Queries)
	genreRepo := postgres.NewGenreRepo(repo.Queries)

	uploadService, err := upload.New(cfg.File.Bucket,
		true, 30*time.Minute,
		store, songRepo, genreRepo, artistRepo)
	if err != nil {
		return nil, err
	}

	// authMw := middleware.NewAuth(nil)
	// rl := ratelimit.NewTokenBucket(15, 30)

	if !testing.Testing() {
		root.Use(middleware.Logger())
		// root.Use(middleware.Limit(nil))

	}
	// oidc, err := auth.NewClient(ctx, "OIDC_ISSUER",
	// 	"OIDC_REDIRECT",
	// 	"OIDC_CLIENT_ID",
	// 	"OIDC_CLIENT_SECRET",
	// )
	// if err != nil {
	// 	return nil, err
	// }
	root.Use(middleware.Cors())

	root.Route("/v1", func(api chi.Router) {
		api.Route("/songs", handleSongs(songRepo, nil, uploadService))
		api.Route("/playlists", handlePlaylists(playlistRepo))
		api.Route("/albums", handleAlbums(albumRepo))

		api.Route("/images", handleImages(uploadService))
		api.Route("/audio", handleAudio(uploadService))
		api.Route("/artists", handleArtists(artistRepo))
		api.Route("/genres", handleGenres(genreRepo))

		// api.Route("/admin", func(r chi.Router) {
		// 	// r.Use(authMw.RequireAuth(auth.RoleAdmin))
		// 	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 	// 	jsonutil.WriteJSON(w, nil, http.StatusOK)
		// 	// })
		// })
		api.Route("/auth", handleAuth())

		api.Route("/users", func(r chi.Router) {
			// r.Get("/", handleUsers(userRepo))

			// r.Get("/{id}", handleGetUsersId(userRepo))
			// r.Delete("/{id}", handleDelteUsersId(userRepo))

			r.Get("/{id}/history", handleUserHistory())
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

// Gets pathvalue {id} from request
func pathValId(r *http.Request) (uuid.UUID, error) {
	return uuid.Parse(r.PathValue("id"))
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
