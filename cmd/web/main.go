package main

import (
	"log"
	"sync"

	"github.com/nathanaelcunningham/billReminder/config"
	"github.com/nathanaelcunningham/billReminder/db"
	"github.com/nathanaelcunningham/billReminder/mailer"
)

type application struct {
	wg         sync.WaitGroup
	db         *db.DB
	billRepo   *db.BillRepository
	mailClient *mailer.Mailer
	cfg        *config.Config
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.NewConfig(".env")
	database := db.New("billReminder.db")
	database.RunMigrations()
	defer database.Close()

	billRepo := db.NewBillRepository(database)

	mailClient := mailer.NewMailer(cfg.SendgridApiKey)

	app := &application{
		wg:         sync.WaitGroup{},
		db:         database,
		billRepo:   billRepo,
		mailClient: mailClient,
		cfg:        cfg,
	}

	return app.serveHTTP()
}
