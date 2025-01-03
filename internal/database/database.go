package database

import "vault/internal/models"

type Client interface {
	GetUniques2() ([]models.UniquesDTO, error)
	GetExch(tableName string) ([]models.ExchDTO, error)
  Health() error
}
