package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	JWTSecret string
	Port      string
	DBPath    string
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port:      os.Getenv("PORT"),
		DBPath:    os.Getenv("DB_PATH"),
	}
	if cfg.JWTSecret == "" {
		log.Fatal().Msg("JWT_SECRET is required")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.DBPath == "" {
		cfg.DBPath = "chinook.db"
	}
	return cfg
}
