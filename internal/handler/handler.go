package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"movieexample.com/internal/auth"
	"movieexample.com/internal/repository"
	"movieexample.com/internal/user"
	"movieexample.com/internal/util/fileutil"
	"movieexample.com/pkg/model"
)

type Handler struct {
	authController *auth.Controller
	userController *user.Controller
}

func New(authController *auth.Controller,
	userController *user.Controller) *Handler {
	return &Handler{
		authController: authController,
		userController: userController,
	}
}

func (h *Handler) Register(mux *http.ServeMux) {
	// setup routes
	mux.HandleFunc("/health", h.handleHealth)
	//mux.Handle("/", middleware.Logger)
	// user routes
	mux.HandleFunc("/user", h.handleUser)

	// auth routes
	mux.HandleFunc("/auth/login", h.handleLogin)
	mux.HandleFunc("/auth/signup", h.handleSignup)
	//mux.HandleFunc("/auth/oath", h.handleSignup)

	//mux.HandleFunc("/secret", h.handleSecret)

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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "empty email or password")
		return
	}
	passwordHash := h.authController.HashPassword(password)
	user := &model.User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	ctx := r.Context()
	if err := h.userController.PutUser(ctx, user); err != nil {

	}

	tokenString, err := h.authController.GenerateToken(&auth.Claims{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, tokenString)
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
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
	_, err := h.userController.GetUserByEmail(ctx, email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO
}

/*
func (h *Handler) handleSecret(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(t *jwt.Token) (any, error) {
		return "", nil
	}, request.WithClaims(&auth.Claims{}))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = fmt.Fprintln(w, "Invalid token: %w", err)
	}
	_, _ = fmt.Fprintln(w, "Welcome,", token.Claims.(*auth.Claims).Username)
}
*/

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
		err := h.userController.PutUser(ctx, &model.User{})
		if err != nil {
			log.Printf("Repository get error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:

		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(http.StatusOK)
}
