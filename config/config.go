package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SendgridApiKey string `mapstructure:"SENDGRID_API_KEY"`
}

func NewConfig(path string) *Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("File .env not found, reading configuration from ENV")
	}

	var config Config
	if config, err = loadConfig(path); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	return &config
}

func loadConfig(path string) (config Config, err error) {
	cfg := Config{
		SendgridApiKey: os.Getenv("SENDGRID_API_KEY"),
	}
	return cfg, nil
}
