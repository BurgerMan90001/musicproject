package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"musicproject.com/config"
	"musicproject.com/internal/middleware"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/util/fileutil"
	"musicproject.com/pkg/util/handleutil"
)

type Handler struct {
	mux  *http.ServeMux
	repo repository.Repository
	cfg  config.Config
}

func New(mux *http.ServeMux, repo repository.Repository, cfg config.Config) *Handler {
	return &Handler{mux, repo, cfg}
}

func (h *Handler) Register(path string) {
	jwtAccessKey := h.cfg.JWTAccessKey()
	oathCfg := h.cfg.GoogleOathConfig()

	// setup routes
	h.mux.HandleFunc("/health", handleHealth)
	// user routes
	h.mux.HandleFunc("/user", handleUser(h.repo))

	h.mux.HandleFunc("/songs", handleSongs(h.repo))
	h.mux.HandleFunc("/artists", handleArtists())

	// auth routes
	h.mux.HandleFunc("/auth/login", handleLogin(jwtAccessKey, h.repo))
	h.mux.HandleFunc("/auth/signup", handleSignup(jwtAccessKey, h.repo))
	h.mux.HandleFunc("/auth/refresh", handleRefresh())
	h.mux.HandleFunc("/auth/email/reset", handleEmailReset())

	h.mux.HandleFunc("/auth/google/login", handleOauthGoogleLogin(oathCfg))
	h.mux.HandleFunc("/auth/google/redirect", handleOauthGoogleRedirect(jwtAccessKey, oathCfg))

	h.mux.HandleFunc("/secret", middleware.JWTMiddleware(jwtAccessKey, handleSecret))

	// static file server
	//mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//middleware.Logger(mux)
}
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
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "alive")
}

func handleSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleSongs(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			song, err := repo.GetSongByID(ctx, id)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				handleutil.InternalServerError(w, r, err)
				return
			}
			fileutil.WriteJSON(w, song)

		case http.MethodPut:
			err := repo.PutSong(ctx, id, nil)
			if err != nil {
				handleutil.InternalServerError(w, r, err)
			}
		default:
			handleutil.MethodNotAllowedError(w, r)
		}
	}
}

func handleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id := r.FormValue("id")
		switch r.Method {
		case http.MethodGet:
			//artist, err
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}
