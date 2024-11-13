package store

import (
	"database/sql"
	"log/slog"
)

// Store is used to facilitate database operations.
type Store struct {
	log *slog.Logger
	db  *sql.DB
}

// NewStateful creates a new Stateful struct.
func New(log *slog.Logger, db *sql.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}
