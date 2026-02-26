package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

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
	authController *auth.Controller
	userController *user.Controller
	cfg            config.ServiceConfig
}

func New(mux *http.ServeMux, cfg config.ServiceConfig) *Handler {

	repo := newRepository(cfg)
	return &Handler{
		mux,
		auth.New([]byte(cfg.APIConfig.JWTKey)),
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
	// setup routes
	h.mux.HandleFunc("/health", h.handleHealth)
	// user routes
	h.mux.HandleFunc("/user", h.handleUser)

	// auth routes
	h.mux.HandleFunc("/auth/login", h.handleLogin)
	h.mux.HandleFunc("/auth/signup", h.handleSignup)
	//mux.HandleFunc("/auth/oath", h.handleSignup)

	h.mux.HandleFunc("/secret", middleware.JWTMiddleware(jwtKey, h.handleSecret))

	// static file server
	//mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//middleware.Logger(mux)
}
func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "alive")
}
func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
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

	tokenString, err := h.authController.GenerateToken(&auth.Claims{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, tokenString)
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPut:
		user := &model.User{}

		err := h.userController.PutUser(ctx, user)
		if err != nil {
			log.Printf("Repository get error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:

		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(http.StatusOK)
}
