package serve

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/thenorthnate/buzz"
	"github.com/thenorthnate/pony/pkg/aol"
)

const (
	defaultAddr = ""
	defaultPort = "8080"
)

func Launch() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := getApp()
	messages := make(chan []byte, 1)
	hive := buzz.New()
	hive.Submit(aol.GetWorker(app.logs, messages))
	s := app.getServer(ctx)
	if err := s.ListenAndServe(); err != nil {
		app.logs.Error("server shutdown with an error", slog.String("err", err.Error()))
	}
}

type app struct {
	logs *slog.Logger
}

func getApp() *app {
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
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/subscribe", a.subscribe)
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
		Handler: mux,
	}
	return s
}
