package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *handler) HealthRoutes() *chi.Mux {
	api := chi.NewRouter()
	api.Get("/health", h.health)
	return api
}

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := h.HealthCheck()
	render.JSON(w, r, response)
}
