package main

import (
	"log"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/nathanaelcunningham/billReminder/assets"
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
	files      *assets.Embed
	uuid       []byte
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	cfg := config.NewConfig(".env")
	database := db.New(cfg.DBPath, true)
	// database.RunMigrations()
	defer database.Close()

	billRepo := db.NewBillRepository(database)

	mailClient := mailer.NewMailer(cfg.SendgridApiKey)

	files := assets.New()
	id := []byte(uuid.New().String())

	app := &application{
		wg:         sync.WaitGroup{},
		db:         database,
		billRepo:   billRepo,
		mailClient: mailClient,
		cfg:        cfg,
		files:      files,
		uuid:       id,
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("00:01").Do(app.RunEmailCron)
	s.StartAsync()

	return app.serveHTTP()
}
