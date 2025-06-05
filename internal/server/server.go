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
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           handler,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	return server
}
