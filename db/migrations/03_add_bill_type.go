package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func init() {
	goose.AddMigrationContext(Up0003, Down0003)
}

func Up0003(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(
		ctx,
		"ALTER TABLE bills ADD COLUMN bill_type TEXT DEFAULT 'STATIC';",
	)
	return err
}

func Down0003(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE bills DROP COLUMN bill_type;")
	return err
}
