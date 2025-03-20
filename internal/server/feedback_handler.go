package server

import (
	"log/slog"
	"net/http"
	"vault/internal/models"
	"vault/internal/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func (s *Server) FeedbackHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		ctx, span := tracer.Start(ctx, "POST feedback", trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		var feedback models.Feedback
		if err := utils.Decode(r, &feedback); err != nil {
			span.SetStatus(codes.Error, "failed to decode feedback")
			span.RecordError(err)

			utils.Error(w, http.StatusBadRequest, "Invalid request body", err, nil)
			return
		}

		problems := feedback.Validate()
		if len(problems) > 0 {
			span.SetStatus(codes.Error, "failed validation")

			utils.Error(
				w,
				http.StatusUnprocessableEntity,
				"Feedback validation failed.",
				nil,
				problems,
			)
			return
		}

		err := s.db.SaveFeedback(ctx, feedback)
		if err != nil {
			span.SetStatus(codes.Error, "failed to save feedback")
			span.RecordError(err)

			utils.Error(
				w,
				http.StatusInternalServerError,
				"An internal server error occurred. Please try again later.",
				nil,
				nil,
			)
			slog.Error("creating strategy", "error", err, "message", feedback)
			return
		}

		span.SetStatus(codes.Ok, "successfully saved the feedback")
		w.WriteHeader(http.StatusOK)
	})
}
