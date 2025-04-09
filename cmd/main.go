package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vault/internal/database"
	"vault/internal/server"
	"vault/pkg/telemetry"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

var (
	service = os.Getenv("SERVICE_NAME")
	logger  = otelslog.NewLogger(service)
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to start api server", "error", err)
		os.Exit(1)
	}
}

func run() error {
	port := os.Getenv("PORT")
	primarURL := os.Getenv("DB_URL")
	authToken := os.Getenv("DB_AUTH_TOKEN")
	cors := os.Getenv("CORS")

	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	otelShutdown, err := telemetry.SetupOTelSDK(ctx)
	if err != nil {
		return err // TODO: add context to error
	}
	defer otelShutdown(context.Background())

	db, err := database.NewReplica(primarURL, authToken)
	if err != nil {
		return err // TODO: add more context
	}

	client := &database.LibsqlClient{DB: db}
	srv := server.New(client, port, cors)

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case err = <-srvErr:
		return err // TODO: more context
	case <-ctx.Done():
		stop()
	}

	ctxTO, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	return srv.Shutdown(ctxTO)
}
