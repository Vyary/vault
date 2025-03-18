package server

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"vault/internal/utils"

	"go.opentelemetry.io/otel"
)

var (
	service = os.Getenv("SERVICE_NAME")
	tracer  = otel.Tracer(service)
)

func (s *Server) Uniques2Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "GET uniques2")
		defer span.End()

		u, err := s.db.GetUniques2(ctx)
		if err != nil {
			utils.Error(
				w,
				http.StatusInternalServerError,
				"An internal server error occurred. Please try again later.",
				err,
				nil,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		if err := json.NewEncoder(gz).Encode(u); err != nil {
			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}
	})
}

func (s *Server) Exch2Handler(tableName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), fmt.Sprintf("GET %s", tableName))
		defer span.End()

		u, err := s.db.GetExch(ctx, tableName)
		if err != nil {
			utils.Error(
				w,
				http.StatusInternalServerError,
				"An internal server error occurred. Please try again later.",
				err,
				nil,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		if err := json.NewEncoder(gz).Encode(u); err != nil {
			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}
	})
}
