// Package db holds simple SQLite helpers.
package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Open opens a SQLite DB with sane defaults (WAL, FKs).
func Open(path string) (*sql.DB, error) {
	if path == "" {
		return nil, fmt.Errorf("empty db path")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"file:%s?mode=rwc&_busy_timeout=%d&_journal_mode=WAL&_foreign_keys=on",
		path,
		2000,
	)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(2)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
