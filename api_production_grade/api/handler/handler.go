package handler

import (
	"github.com/arkiant/golab/api_production_grade/api/service"
)

// Handler type represents all api core endpoints
type Handler struct {
	*service.Service
}

// NewHandler create a new handler type
func NewHandler(s service.Service) *Handler {
	return &Handler{&s}
}
