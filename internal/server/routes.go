package server

import (
	"net/http"
	"vault/internal/middleware"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /health", s.HealthHandler())

	mux.Handle("GET /v1/uniques2", middleware.CacheControl(s.Uniques2Handler(), 5))
	mux.Handle("GET /v1/fragments2", middleware.CacheControl(s.Exch2Handler("fragments2"), 5))
	mux.Handle("GET /v1/runes2", middleware.CacheControl(s.Exch2Handler("runes2"), 5))
	mux.Handle("GET /v1/essences2", middleware.CacheControl(s.Exch2Handler("essences2"), 5))
	mux.Handle("GET /v1/cores2", middleware.CacheControl(s.Exch2Handler("cores2"), 5))
	mux.Handle("GET /v1/catalysts2", middleware.CacheControl(s.Exch2Handler("catalysts2"), 5))
	mux.Handle("GET /v1/artifacts2", middleware.CacheControl(s.Exch2Handler("expedition2"), 5))
	mux.Handle("GET /v1/omens2", middleware.CacheControl(s.Exch2Handler("omens2"), 5))
	mux.Handle("GET /v1/distillations2", middleware.CacheControl(s.Exch2Handler("distillations2"), 5))
	mux.Handle("GET /v1/waystones2", middleware.CacheControl(s.Exch2Handler("waystones2"), 5))

	return mux
}
