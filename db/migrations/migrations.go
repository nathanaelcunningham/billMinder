package migrations

import (
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func Run(dbPath string) error {
	db, err := goose.OpenDBWithDriver("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	err = goose.Run("up", db, ".")
	if err != nil {
		return err
	}
	return nil
}
