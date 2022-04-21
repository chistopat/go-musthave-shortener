package shortener

import (
	"net/http"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", "https://ya.ru")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func ShortURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
