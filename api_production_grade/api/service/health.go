package service

import (
	"github.com/arkiant/golab/api_production_grade/models"
)

// HealthCheck Check health API
func (s *Service) HealthCheck() models.Health {
	return models.Health{Message: "OK"}
}
