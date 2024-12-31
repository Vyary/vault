package server

import (
	"net/http"
	"vault/internal/middleware"
	"vault/ui"
)

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(ui.Web))

	mux.Handle("GET /", middleware.CacheControl(fileServer, 24*60))
	mux.Handle("GET /v1/uniques2", middleware.CacheControl(s.GetUniques2(), 5))
	mux.Handle("GET /v1/runes2", middleware.CacheControl(s.GetRunes2(), 5))
	mux.Handle("GET /v1/cores2", middleware.CacheControl(s.GetCores2(), 5))
	mux.Handle("GET /v1/fragments2", middleware.CacheControl(s.GetFragments2(), 5))

	return mux
}
