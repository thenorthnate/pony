package store

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type Message struct {
	ID          uint64
	MessageID   string
	Source      string
	ContentType string
	Data        []byte
	ReceivedAt  time.Time
}

// InsertMessage inserts a new message to the given table.
func (s *Store) InsertMessage(
	ctx context.Context,
	topic string,
	contentType string,
	data []byte,
	queryValues url.Values,
) error {
	msg := Message{
		MessageID:   queryValues.Get("message_id"),
		Source:      queryValues.Get("source"),
		ContentType: contentType,
		Data:        data,
		ReceivedAt:  time.Now().UTC(),
	}
	s.log.Info(fmt.Sprintf("Got Message: %v", msg))
	return nil
}
