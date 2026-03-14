package handleutil

import (
	"fmt"
	"log"
	"net/http"
)

func MethodNotAllowedError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("%v method not allowed", r.Method), http.StatusMethodNotAllowed)
}

func NotFoundError(w http.ResponseWriter, r *http.Request) {

}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %v", err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
