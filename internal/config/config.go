// Package config for configuration and settings
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("Env"),
	}
	return cfg, nil
}
