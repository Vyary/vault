package database

import (
	"context"
	"fmt"
	"vault/internal/models"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (l *LibsqlClient) SaveFeedback(ctx context.Context, feedback models.Feedback) error {
	_, span := tracer.Start(ctx, "INSERT feedback", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	query := `
  INSERT INTO feedback (name, email, message)
  VALUES (?, ?, ?)`

	_, err := l.DB.Exec(query, &feedback.Name, &feedback.Email, &feedback.Message)
	if err != nil {
		span.SetStatus(codes.Error, "failed to save feedback")
		span.RecordError(err)

		return fmt.Errorf("failed to save a feedback, %w", err)
	}

	span.SetStatus(codes.Ok, "successfully saved feedback")
	return nil
}
