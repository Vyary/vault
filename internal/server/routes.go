package server

import (
	"net/http"
	"vault/internal/middleware"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /health", s.HealthHandler())
	mux.Handle("POST /contact", s.FeedbackHandler())

	mux.Handle("GET /v2/{league}/uniques2", middleware.CacheControl(s.UniquesHandler2(), 5))
	mux.Handle("GET /v2/{league}/{table}", middleware.CacheControl(s.ExchHandler2(), 5))

	return mux
}
