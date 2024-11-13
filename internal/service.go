package pony

import (
	"context"
	"log/slog"

	"github.com/thenorthnate/buzz"
	"github.com/thenorthnate/pony/internal/pkg/aol"
)

const (
	defaultAddr = ""
	defaultPort = "8080"
)

func Launch() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := newApp()
	messages := make(chan []byte, 1)
	hive := buzz.New()
	hive.Submit(aol.GetWorker(app.logs, messages))
	s := app.getServer(ctx)
	if err := s.ListenAndServe(); err != nil {
		app.logs.Error("server shutdown with an error", slog.String("err", err.Error()))
	}
}
