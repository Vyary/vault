package database

import "vault/internal/models"

type Client interface {
  GetUniques2() ([]models.UniquesDTO, error)
}
