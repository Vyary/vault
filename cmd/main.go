package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vault/internal/database"
	"vault/internal/proxy"
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
	primaryUrl := os.Getenv("DB_URL")
	authToken := os.Getenv("DB_AUTH_TOKEN")
	proxyTarget := os.Getenv("PROXY_TARGET")
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := database.NewRemote(primaryUrl, authToken)
	if err != nil {
		slog.Error("Error initializing database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	client := &database.LibsqlClient{DB: db}
	server := server.New(client, port)
	done := make(chan struct{})

	proxy, err := proxy.New(proxyTarget)
	if err != nil {
		slog.Error("starting proxy", "error", err)
	}

	go gracefulShutdown(server, done)

	go func() {
		slog.Info(fmt.Sprintf("Starting API Server on %s...", port))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("HTTP server error", "error", err)
		}
	}()

	go func() {
		slog.Info("Proxy Server on :443...")

		if err := proxy.ListenAndServeTLS(certFile, keyFile); err != nil {
			slog.Error("Proxy server error", "error", err)
		}
	}()

	<-done
	slog.Info("Graceful shutdown complete.")
}
