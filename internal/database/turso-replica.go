package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tursodatabase/go-libsql"
)

type LibsqlClient struct {
	DB *sql.DB
}

func NewReplica(primaryUrl, authToken string) (*sql.DB, error) {
	dir := "./internal/database/local"
	dbName := "local.db"

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating directory '%s': %w", dir, err)
	}

	dbPath := filepath.Join(dir, dbName)
	syncInterval := time.Minute

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
		libsql.WithSyncInterval(syncInterval),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating connector: %w", err)
	}

	db := sql.OpenDB(connector)
	return db, nil
}

func (l *LibsqlClient) Health() error {
  return l.DB.Ping()
}
