package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"vault/internal/database"
	"vault/internal/server"
)

func gracefulShutdown(apiServer *http.Server, done chan struct{}) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	slog.Info("Shutting down gracefully, press Ctrl+C again to force")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server forced to shutdown with error", "error", err)
	}

	slog.Info("Server exiting")
	close(done)
}

func main() {
	port := os.Getenv("PORT")
	primaryUrl := os.Getenv("TURSO_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := database.NewReplica(primaryUrl, authToken)
	if err != nil {
		slog.Error("Error initializing database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
  
	client := &database.LibsqlClient{DB: db}

	server := server.New(client, port)
	done := make(chan struct{})

	go gracefulShutdown(server, done)

	slog.Info("Server is running", "port", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("HTTP server error", "error", err)
	}

	<-done
	slog.Info("Graceful shutdown complete.")
}
