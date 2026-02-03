package migrator

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Migrator struct {
	db            *sql.DB
	migrationsDir string
}

func NewMigrator(pool *pgxpool.Pool, migrationsDir string) *Migrator {
	db := stdlib.OpenDBFromPool(pool)

	return &Migrator{
		db:            db,
		migrationsDir: migrationsDir,
	}
}

func (m *Migrator) Up() error {
	err := goose.Up(m.db, m.migrationsDir)
	if err != nil {
		return err
	}

	return nil
}
