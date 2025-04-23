package rest

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func route(r *chi.Mux) {
	r.Get("/hello", getHello)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "hello")
}
