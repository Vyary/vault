package database

import (
	"database/sql"
	"fmt"
)

func NewRemote(primaryUrl, authToken string) (*sql.DB, error) {
	url := fmt.Sprintf("%s?authToken=%s", primaryUrl, authToken)

	return sql.Open("libsql", url)
}
