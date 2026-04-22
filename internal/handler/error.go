package handler

import (
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/pkg/model"
)


func handleNotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonutil.WriteError(w, &model.Error{
			Code:    http.StatusNotFound,
			Message: "Route not found",
		})
	})
}
func handleNotImplemented() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonutil.WriteError(w, &model.Error{
			Code:    http.StatusNotImplemented,
			Message: "Not implemented",
		})
	})
}
