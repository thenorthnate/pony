package pony

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

type app struct {
	logs *slog.Logger
	db   Database
}

func newApp() *app {
	return &app{
		logs: getLogger(),
	}
}

func getLogger() *slog.Logger {
	var out io.Writer = os.Stderr
	destination := os.Getenv("LOGS_DESTINATION")
	if destination == "STDOUT" {
		out = os.Stdout
	}
	levelStr := os.Getenv("LOG_LEVEL")
	level := slog.LevelInfo
	switch levelStr {
	case slog.LevelDebug.String():
		level = slog.LevelDebug
	}
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	textLogs, _ := strconv.ParseBool(os.Getenv("TEXT_LOGS"))
	if textLogs {
		return slog.New(slog.NewTextHandler(out, options))
	}
	return slog.New(slog.NewJSONHandler(out, options))
}

func (a *app) getServer(ctx context.Context) *http.Server {
	ponyAddr := os.Getenv("PONY_ADDR")
	if ponyAddr == "" {
		ponyAddr = defaultAddr
	}
	ponyPort := os.Getenv("PONY_PORT")
	if ponyPort == "" {
		ponyPort = defaultPort
	}
	a.logs.Info(
		"starting pony server",
		slog.String("pony_addr", ponyAddr),
		slog.String("pony_port", ponyPort),
	)
	s := &http.Server{
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
		Addr:    net.JoinHostPort(ponyAddr, ponyPort),
		Handler: a.createViews(),
	}
	return s
}
