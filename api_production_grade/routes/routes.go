package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Health status return OK
func HealthRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := health{Message: "OK"}
		render.JSON(w, r, response)
	})

	return r
}

type health struct {
	Message string `json:"message"`
}