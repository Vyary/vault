package database

import (
	"database/sql"
	"vault/internal/models"
)

type Client interface {
	GetUniques2() ([]models.UniquesDTO, error)
	GetExch(tableName string) ([]models.ExchDTO, error)
	Health() error
}

type LibsqlClient struct {
	DB *sql.DB
}

func (l *LibsqlClient) Health() error {
  return l.DB.Ping()
}
