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
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrator: failed to set dialect: %w", err)
	}

	err = goose.Up(m.db, "migrations")
	if err != nil {
		return fmt.Errorf("migrator: failed to up migrations: %w", err)
	}

	return nil
}

func (m *migrator) Close() error {
	return m.db.Close()
}
