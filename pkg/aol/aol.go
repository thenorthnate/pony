package aol

import (
	"context"
	"log/slog"

	"github.com/thenorthnate/buzz"
	"github.com/thenorthnate/evs"
)

// AOL implements an append-only-log.
type AOL struct {
	logs     *slog.Logger
	messages chan []byte
}

func GetWorker(logs *slog.Logger, messages chan []byte) *buzz.Worker {
	l := &AOL{
		logs:     logs,
		messages: messages,
	}
	return buzz.NewWorker(l)
}

func (l *AOL) Do(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return evs.From(ctx.Err()).Err()
	case msg := <-l.messages:
		l.logs.Info("received new message", slog.String("message", string(msg)))
	}
	return nil
}
