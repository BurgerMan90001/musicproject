package handler

import (
	"net/http"
)

type Cookie interface {
	Cookie(value string, maxAge int) *http.Cookie
	Clear() *http.Cookie
}

func setCookie(w http.ResponseWriter, cookie Cookie, value string, maxAge int) {
	http.SetCookie(w, cookie.Cookie(value, maxAge))
}
func clearCookie(w http.ResponseWriter, cookie Cookie) {
	http.SetCookie(w, cookie.Clear())
}
