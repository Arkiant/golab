package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Health status return OK
func Health() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
