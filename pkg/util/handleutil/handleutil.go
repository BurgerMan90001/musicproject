package handleutil

import (
	"fmt"
	"net/http"
)

func ErrMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("%v method not allowed", r.Method), http.StatusMethodNotAllowed)
}

func ErrNotFound(w http.ResponseWriter, r *http.Request) {

}
