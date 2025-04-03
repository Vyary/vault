package server

import (
	"fmt"
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

	handler := middleware.Time(s.routes())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
