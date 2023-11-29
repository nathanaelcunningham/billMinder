package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func init() {
	goose.AddMigrationContext(Up0004, Down0004)
}

func Up0004(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(
		ctx,
		"CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT);",
	)
	return err
}

func Down0004(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS users;")
	return err
}
