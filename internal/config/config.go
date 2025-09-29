// Package config for configuration and settings
package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	Env    string
	DBPath string
}

func LoadConfig() (*Config, error) {
	slog.Info("Loading Config")
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:   os.Getenv("PORT"),
		Env:    os.Getenv("ENV"),
		DBPath: os.Getenv("DB_PATH"),
	}
	return cfg, nil
}
