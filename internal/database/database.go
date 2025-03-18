package database

import (
	"context"
	"database/sql"
	"vault/internal/models"
)

type Client interface {
	GetUniques2(ctx context.Context) ([]models.UniquesDTO, error)
	GetExch(ctx context.Context, tableName string) ([]models.ExchDTO, error)
	Health() error
	SaveFeedback(ctx context.Context, feedback models.Feedback) error
}

type LibsqlClient struct {
	DB *sql.DB
}

func (l *LibsqlClient) Health() error {
	return l.DB.Ping()
}
