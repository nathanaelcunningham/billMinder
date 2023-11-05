package main

import (
	"log"
	"sync"

	"github.com/nathanaelcunningham/billReminder/db"
)

type application struct {
	wg       sync.WaitGroup
	db       *db.DB
	billRepo *db.BillRepository
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	database := db.New("billReminder.db")
	database.RunMigrations()
	defer database.Close()

	billRepo := db.NewBillRepository(database)

	app := &application{
		wg:       sync.WaitGroup{},
		db:       database,
		billRepo: billRepo,
	}
	return app.serveHTTP()
}
