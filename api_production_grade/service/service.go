package service

import (
	"database/sql"
)

// Service is a core type and dependency injection
// this type has a business logic
type Service struct {
	db      *sql.DB
	version string
}

// NewService create a new service
func NewService(db *sql.DB, version string) *Service {
	return &Service{
		db:      db,
		version: version,
	}
}

// Version configured api
func (s *Service) Version() string {
	return s.version
}
