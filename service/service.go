package service

import (
	"github.com/jmoiron/sqlx"
)

// Service contains core logic used to back Rest API
type Service struct {
	db *sqlx.DB
	origin string
}

// New Service implementation 
func New(db *sqlx.DB, origin string) *Service {
	return &Service{db: db, origin: origin}
}