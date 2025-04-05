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

func (s *Server) UniquesHandler2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		league := r.PathValue("league")

		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		ctx, span := tracer.Start(ctx,
			fmt.Sprintf("GET %s uniques2", league),
			trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		u, err := s.db.GetUniques2(ctx, league)
		if err != nil {
			span.SetStatus(codes.Error, fmt.Sprintf("failed to retrieve %s uniques2", league))
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
			span.SetStatus(codes.Error, fmt.Sprintf("failed to encode %s uniques2", league))
			span.RecordError(err)

			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}

		span.SetStatus(
			codes.Ok,
			fmt.Sprintf("successfully retrieved and returned %s uniques2", league),
		)
	})
}

func (s *Server) ExchHandler2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		league := r.PathValue("league")
		table := r.PathValue("table")

		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		ctx, span := tracer.Start(
			ctx,
			fmt.Sprintf("GET %s %s", league, table),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		u, err := s.db.GetExch(ctx, league, table)
		if err != nil {
			span.SetStatus(codes.Error, fmt.Sprintf("failed to retrieve %s %s", league, table))
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
			span.SetStatus(codes.Error, fmt.Sprintf("failed to encode %s %s", league, table))
			span.RecordError(err)

			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}

		span.SetStatus(
			codes.Ok,
			fmt.Sprintf("successfully retrieved and returned %s %s", league, table),
		)
	})
}
