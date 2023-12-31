package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/nathanaelcunningham/billReminder/db/migrations"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func New(path string, automigrate bool) *DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal(err)
	}

	if automigrate {
		err = migrations.Run(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &DB{db}
}

func (db *DB) RunMigrations() error {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS bills (id INTEGER PRIMARY KEY, name TEXT, due_date_day INTEGER, amount REAL)",
	)
	return err
}

func Teardown(dbName string) {
	rootDir, _ := os.Getwd()
	err := os.Remove(filepath.Join(rootDir, dbName))
	if err != nil {
		log.Fatal(err)
	}
}
