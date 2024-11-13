package store

import (
	"context"
	"time"
)

type Message struct {
	ID        uint64
	MessageID string
	Source    string
	Data      []byte
	CreatedAt time.Time
}

// InsertMessage inserts a new message to the given table.
func (s *Store) InsertMessage(ctx context.Context, tableName string, m *Message) error {
	return nil
}
