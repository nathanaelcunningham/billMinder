package db

import (
	"os"
	"testing"
)

var database *DB

func setup() {
	database = New("test.db")
	database.RunMigrations()
}

func teardown() {
	Teardown("test.db")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
