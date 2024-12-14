package server

import (
	"net/http"
	"vault/internal/middleware"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("GET /api/v1/poe2/uniques", middleware.CacheControl(s.GetUniques2()))

	return mux
}
