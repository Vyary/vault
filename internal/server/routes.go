package server

import (
	"net/http"
	_ "net/http/pprof"
	"vault/internal/middleware"
	"vault/ui"
)

func (s *Server) routes() http.Handler {
	staticFiles := http.FileServer(http.FS(ui.Web))
	indexFile := http.FileServer(http.FS(ui.Index))

	cacheIndex := middleware.CacheControl(indexFile, 24*60)
	cacheFiles := middleware.CacheControl(staticFiles, 24*60)

	mux := http.NewServeMux()

	mux.Handle("GET /", cacheFiles)
	mux.Handle("GET /uniques", http.StripPrefix("/uniques", cacheIndex))
	mux.Handle("GET /fragments", http.StripPrefix("/fragments", cacheIndex))
	mux.Handle("GET /runes", http.StripPrefix("/runes", cacheIndex))
	mux.Handle("GET /essences", http.StripPrefix("/essences", cacheIndex))
	mux.Handle("GET /cores", http.StripPrefix("/cores", cacheIndex))
	mux.Handle("GET /catalysts", http.StripPrefix("/catalysts", cacheIndex))
	mux.Handle("GET /artifacts", http.StripPrefix("/artifacts", cacheIndex))
	mux.Handle("GET /omens", http.StripPrefix("/omens", cacheIndex))
	mux.Handle("GET /distillations", http.StripPrefix("/distillations", cacheIndex))
	mux.Handle("GET /tablets", http.StripPrefix("/tablets", cacheIndex))

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
