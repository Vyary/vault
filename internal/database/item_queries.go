package database

import (
	"fmt"
	"vault/internal/models"
)

func (l *LibsqlClient) GetUniques2() ([]models.UniquesDTO, error) {
	query := `
  WITH latest_prices AS (
    SELECT 
      item_id, 
      value, 
      type, 
      listed
    FROM prices
    GROUP BY item_id
    HAVING timestamp = MAX(timestamp)
  )
  SELECT 
    u.item_id, 
    u.name, 
    u.base, 
    COALESCE(i.image, '') AS image, 
    COALESCE(lp.value, 0) AS value, 
    COALESCE(lp.type, '') AS type, 
    COALESCE(lp.listed, 0) AS listed
  FROM 
    uniques2 u
    LEFT JOIN images i ON u.item_id = i.item_id
    LEFT JOIN latest_prices lp ON u.item_id = lp.item_id`

	rows, err := l.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query uniques items for poe2: %w", err)
	}
	defer rows.Close()

	var uniques []models.UniquesDTO

	for rows.Next() {
		var unique models.UniquesDTO

		err := rows.Scan(&unique.ItemID, &unique.Name, &unique.Base, &unique.Image, &unique.Price.Value, &unique.Price.Type, &unique.Price.Listed)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poe2 unique: %w", err)
		}

		uniques = append(uniques, unique)
	}

	return uniques, nil
}

func (l *LibsqlClient) GetExch(tableName string) ([]models.ExchDTO, error) {
	query := fmt.Sprintf(`
  WITH latest_prices AS (
    SELECT 
      item_id, 
      value, 
      type, 
      listed
    FROM prices
    GROUP BY item_id
    HAVING timestamp = MAX(timestamp)
  )
  SELECT 
    u.item_id, 
    u.name, 
    COALESCE(i.image, '') AS image, 
    COALESCE(lp.value, 0) AS value, 
    COALESCE(lp.type, '') AS type, 
    COALESCE(lp.listed, 0) AS listed
  FROM %s u
  LEFT JOIN images i ON u.item_id = i.item_id
  LEFT JOIN latest_prices lp ON u.item_id = lp.item_id`, tableName)

	rows, err := l.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query %s for poe2: %w", tableName, err)
	}
	defer rows.Close()

	var exchItems []models.ExchDTO

	for rows.Next() {
		var exchItem models.ExchDTO

		err := rows.Scan(&exchItem.ItemID, &exchItem.Name, &exchItem.Image, &exchItem.Price.Value, &exchItem.Price.Type, &exchItem.Price.Listed)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poe2 %s: %w", tableName, err)
		}

		exchItems = append(exchItems, exchItem)
	}

	return exchItems, nil
}
