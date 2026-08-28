package config

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewDatabase(DatabaseConfig DatabaseConfig) *sqlx.DB {
	if DatabaseConfig.Driver == "" {
		log.Fatal("please fill in database credentials in .env file or set in environment variable")
	}

	db, err := sqlx.Open(DatabaseConfig.Driver, DatabaseConfig.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	// Set the below later
	// s.Db.SetMaxOpenConns(10)
	// s.Db.SetMaxIdleConns(10)
	// s.Db.SetConnMaxLifetime(10)

	return db
}
