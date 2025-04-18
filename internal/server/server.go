package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"vault/internal/database"
	"vault/internal/middleware"
)

type Server struct {
	db database.Client
}

func New(db database.Client, port string, cors string) *http.Server {
	s := &Server{db: db}

	handler := middleware.Time(s.routes())

	if cors == "TRUE" {
		handler = middleware.Cors(handler)
		slog.Info("CORS enabled")
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
