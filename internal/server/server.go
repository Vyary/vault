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

func New(db database.Client, port string) *http.Server {
	s := &Server{db: db}

	handler := middleware.Cors(middleware.Time(s.routes()))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}


	return server
}
