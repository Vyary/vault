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

	handler := middleware.Cors(middleware.Time(s.routes()))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
