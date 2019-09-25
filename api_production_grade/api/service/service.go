package service

import (
	"database/sql"
)

// Service is a core type and dependency injection
// this type has a business logic
type Service struct {
	db *sql.DB
}

// NewService create a new service
func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}
