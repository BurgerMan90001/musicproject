package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"okapi.com/config"
	"okapi.com/internal/auth"
	"okapi.com/internal/middleware"
	"okapi.com/internal/repository"
	"okapi.com/internal/repository/memory"
	"okapi.com/internal/repository/postgres"
	"okapi.com/internal/user"
	"okapi.com/internal/util/fileutil"
	"okapi.com/internal/util/uuid"
	"okapi.com/pkg/model"
)

type Handler struct {
	mux            *http.ServeMux
	userController *user.Controller
	cfg            config.ServiceConfig
}

func New(mux *http.ServeMux, cfg config.ServiceConfig) *Handler {

	repo := newRepository(cfg)
	return &Handler{
		mux,
		user.New(repo),
		cfg,
	}
}
func newRepository(cfg config.ServiceConfig) repository.Repository {
	var repo repository.Repository

	switch cfg.RepositoryConfig.Type {
	case "memory":
		repo = memory.New()
	case "postgres":
		repo = postgres.New(cfg.RepositoryConfig.URL)
	default:
		repo = memory.New()
	}
	return repo
}

func (h *Handler) Register(path string) {
	jwtKey := []byte(h.cfg.APIConfig.JWTKey)
	oathCfg := &oauth2.Config{
		ClientID:     "666665152595-2frtq2bbppq3cb83u5rm7p36hcsgtips.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8081/auth/oauth/redirect",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	// setup routes
	h.mux.HandleFunc("/health", h.handleHealth)
	// user routes
	h.mux.HandleFunc("/user", h.handleUser)

	// auth routes
	h.mux.HandleFunc("/auth/login", h.handleLogin(jwtKey))
	h.mux.HandleFunc("/auth/signup", h.handleSignup(jwtKey))

	h.mux.HandleFunc("/auth/oauth", h.handleOath(oathCfg))
	h.mux.HandleFunc("/auth/oauth/redirect", h.handleOathRedirect(oathCfg))

	h.mux.HandleFunc("/secret", middleware.JWTMiddleware(jwtKey, h.handleSecret))

	// static file server
	//mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//middleware.Logger(mux)
}
func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "alive")
}
func (h *Handler) handleSignup(jwtKey []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "empty email or password")
			return
		}
		passwordHash := auth.HashPassword(password)
		user := &model.User{
			ID:           uuid.GenerateID(),
			Email:        email,
			PasswordHash: passwordHash,
		}
		ctx := r.Context()
		if err := h.userController.PutUser(ctx, user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "repository put error:", err)
			return
		}

		tokenString, err := auth.GenerateToken(jwtKey, &auth.Claims{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/jwt")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, tokenString)
	}
}
func (h *Handler) handleOath(cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	}
}
func (h *Handler) handleOathRedirect(cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		token, err := cfg.Exchange(context.Background(), code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("unable to get token: %v", err)
		}
		fmt.Fprintf(w, "Access Token: %s", token.AccessToken)
	}
}
func (h *Handler) handleLogin(jwtKey []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "empty username or password")
			return
		}
		user, err := h.userController.GetUserByEmail(ctx, email)

		if err != nil && errors.Is(err, repository.ErrNotFound) ||
			!auth.CheckPasswordHash(password, user.PasswordHash) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("invalid email or password")
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println("Login successful")
	}
}

func (h *Handler) handleSecret(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome,", token.Claims.(*auth.Claims).Username)
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) handleUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	switch r.Method {
	case http.MethodGet:
		user, err := h.userController.GetUserByID(ctx, id)
		if err != nil && errors.Is(err, repository.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Repository get error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fileutil.WriteJSON(w, user)

	case http.MethodDelete:
		/*
			claims, err := auth.JWTParseToken(jwtKey, r)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Invalid token:", err)
				return
			}
		*/

	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(http.StatusOK)
}
