package handler

import (
	"github.com/arkiant/golab/api_production_grade/api/service"
	"github.com/go-chi/chi"
)

// handler type represents all api core endpoints
type handler struct {
	*service.Service
}

// NewHandler create a new handler type
func NewHandler(s *service.Service) *chi.Mux {

	h := &handler{s}

	// Module routes
	api := chi.NewRouter()
	api.Mount("/status", h.HealthRoutes())

	// Base api convention
	r := chi.NewRouter()
	r.Route("/v"+h.Version(), func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Mount("/", api)
		})
	})

	return r
}
