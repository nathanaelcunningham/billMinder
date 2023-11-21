package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func init() {
	goose.AddMigrationContext(Up0001, Down0001)
}

func Up0001(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(
		ctx,
		"CREATE TABLE IF NOT EXISTS bills (id INTEGER PRIMARY KEY, name TEXT, due_date_day INTEGER, amount REAL);",
	)
	return err
}

func Down0001(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS bills;")
	return err
}
