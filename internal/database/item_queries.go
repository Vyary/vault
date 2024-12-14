package database

import (
	"fmt"

	"vault/internal/models"
)

func (l *LibsqlClient) GetUniques2() ([]models.UniquesDTO, error) {
	query := `
  SELECT u.name, u.base, COALESCE(i.image, ''), COALESCE(p.value, 0), COALESCE(p.type, ''), COALESCE(p.listed, 0)
  FROM uniques2 u
  LEFT JOIN images i ON u.item_id = i.item_id
  LEFT JOIN prices p ON u.item_id = p.item_id`

	rows, err := l.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query uniques items for poe2: %w", err)
	}

	var uniques []models.UniquesDTO

	for rows.Next() {
		var unique models.UniquesDTO

		err := rows.Scan(&unique.Name, &unique.Base, &unique.Image, &unique.Price.Value, &unique.Price.Type, &unique.Price.Listed)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poe2 unique: %w", err)
		}

		uniques = append(uniques, unique)
	}

	return uniques, nil
}
