package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func init() {
	goose.AddMigrationContext(Up0002, Down0002)
}

func Up0002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(
		ctx,
		"ALTER TABLE bills ADD COLUMN is_autopay INTEGER DEFAULT 0;",
	)
	return err
}

func Down0002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE bills DROP COLUMN is_autopay;")
	return err
}
