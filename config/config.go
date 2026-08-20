package config

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Db *sql.DB
}

func GetConfigurations() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading .env file.")
	}

	cfg := &Config{
		Db: GetDBConnection(),
	}

	return cfg
}
