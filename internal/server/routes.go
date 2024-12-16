package server

import (
	"net/http"
	"vault/internal/middleware"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("GET /api/v1/uniques2", middleware.CacheControl(s.GetUniques2()))

	return mux
}
