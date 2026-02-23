package handler

import (
	"net/http"

	"movieexample.com/internal/controller/auth"
	"movieexample.com/internal/model"
	"movieexample.com/internal/util/fileutil"
)

type Handler struct {
	authController auth.Controller
}

func New(authController auth.Controller) *Handler {
	return &Handler{
		authController: authController,
	}
}
func (h *Handler) Register(handler *http.ServeMux) {
	mux := http.NewServeMux()

	// setup routes
	//mux.HandleFunc("GET /health", handler.HandleHealth)

	// user routes
	mux.HandleFunc("/user", h.handleUser)

	// auth routes
	mux.HandleFunc("POST /auth/login", h.handleLogin)
	mux.HandleFunc("/auth/signup", h.handleSignup)

	mux.HandleFunc("/secret", h.handleSecret)

	// static file server
	mux.Handle("/static/", http.FileServer(http.Dir("public")))

	// add middleware
	//httphandler := middleware.Logger(mux)
	//httphandler = middleware.PanicRecovery(mux)
}
func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	fileutil.WriteJSON(w, &model.User{ID: "asda"})
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleSecret(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleUser(w http.ResponseWriter, r *http.Request) {

}
