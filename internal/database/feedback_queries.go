package database

import (
	"fmt"
	"vault/internal/models"
)

func (l *LibsqlClient) SaveFeedback(feedback models.Feedback) error {
	query := `
  INSERT INTO feedback (name, email, message)
  VALUES (?, ?, ?)`

	_, err := l.DB.Exec(query, &feedback.Name, &feedback.Email, &feedback.Message)
	if err != nil {
		return fmt.Errorf("failed to save a feedback, %w", err)
	}

	return nil
}
