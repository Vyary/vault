package server

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"vault/internal/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var (
	service = os.Getenv("SERVICE_NAME")
	tracer  = otel.Tracer(service)
)

func (s *Server) Uniques2Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		ctx, span := tracer.Start(ctx, "GET uniques2", trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		u, err := s.db.GetUniques2(ctx)
		if err != nil {
			span.SetStatus(codes.Error, "failed to retrieve uniques2")
			span.RecordError(err)

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
			span.SetStatus(codes.Error, "failed to encode uniques2")
			span.RecordError(err)

			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}

		span.SetStatus(codes.Ok, "successfully retrieved and returned uniques2")
	})
}

func (s *Server) Exch2Handler(tableName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		ctx, span := tracer.Start(
			ctx,
			fmt.Sprintf("GET %s", tableName),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		u, err := s.db.GetExch(ctx, tableName)
		if err != nil {
			span.SetStatus(codes.Error, fmt.Sprintf("failed to retrieve %s", tableName))
			span.RecordError(err)

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
			span.SetStatus(codes.Error, fmt.Sprintf("failed to encode %s", tableName))
			span.RecordError(err)

			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}

		span.SetStatus(codes.Ok, fmt.Sprintf("successfully retrieved and returned %s", tableName))
	})
}
