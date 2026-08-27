package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// Global variable
var Cfg Config

func Load() (Config, error) {
	// Load .env during local development.
	// In production, environment variables can be provided
	// directly by Docker/Kubernetes/cloud infrastructure.
	_ = godotenv.Load()

	cfg := Config{
		Server: ServerConfig{
			Port: LoadServerPort(),
		},
		Database: LoadDatabase(),
	}

	if cfg.Database.Driver == "" {
		return Config{}, fmt.Errorf("DB_DRIVER is required")
	}

	if cfg.Database.DSN == "" {
		return Config{}, fmt.Errorf("DB_DSN is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func SetConfig(cfg Config) {
	Cfg = cfg
}

func LoadDatabase() DatabaseConfig {
	var Database DatabaseConfig

	Database.Driver = os.Getenv("DB_DRIVER")
	Database.DSN = os.Getenv("DB_DSN")

	return Database
}

func LoadServerPort() string {

	//var Port string
	//Port := getEnv("SERVER_PORT", "8080")
	//return Port

	return getEnv("SERVER_PORT", "8080")
}
