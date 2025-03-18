package database

import (
	"context"
	"fmt"
	"vault/internal/models"
)

func (l *LibsqlClient) SaveFeedback(ctx context.Context, feedback models.Feedback) error {
	ctx, span := tracer.Start(ctx, "INSERT feedback")
	defer span.End()

	query := `
  INSERT INTO feedback (name, email, message)
  VALUES (?, ?, ?)`

	_, err := l.DB.Exec(query, &feedback.Name, &feedback.Email, &feedback.Message)
	if err != nil {
		return fmt.Errorf("failed to save a feedback, %w", err)
	}

	return nil
}
