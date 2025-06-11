package migrator

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
)

type Migrator interface {
	MigrateUp() error
	Close() error
}

type migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) Migrator {
	return &migrator{
		db: db,
	}
}

func (m *migrator) MigrateUp() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("migrator: failed to set dialect: %w", err)
	}

	if err := goose.Up(m.db, "./migrations"); err != nil {
		return fmt.Errorf("migrator: failed to migrate up: %w", err)
	}

	return nil
}

func (m *migrator) Close() error {
	if err := m.db.Close(); err != nil {
		return fmt.Errorf("migrator: failed to close db: %w", err)
	}

	return nil
}
