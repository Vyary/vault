package server

import (
	"net/http"
	"vault/internal/middleware"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /health", s.HealthHandler())

	mux.Handle("GET /v1/uniques2", s.Uniques2Handler())
	mux.Handle("GET /v1/fragments2", s.Exch2Handler("fragments2"))
	mux.Handle("GET /v1/runes2", s.Exch2Handler("runes2"))
	mux.Handle("GET /v1/essences2", s.Exch2Handler("essences2"))
	mux.Handle("GET /v1/cores2", s.Exch2Handler("cores2"))
	mux.Handle("GET /v1/catalysts2", s.Exch2Handler("catalysts2"))
	mux.Handle("GET /v1/artifacts2", s.Exch2Handler("expedition2"))
	mux.Handle("GET /v1/omens2", s.Exch2Handler("omens2"))
	mux.Handle("GET /v1/distillations2", s.Exch2Handler("distillations2"))
	mux.Handle("GET /v1/waystones2", s.Exch2Handler("waystones2"))
	mux.Handle("GET /v1/bases2", s.Exch2Handler("bases2"))

	return middleware.CacheControl(mux, 5)
}
