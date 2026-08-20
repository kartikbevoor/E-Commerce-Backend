package config

import (
	"database/sql"
	"log"
	"os"
)

func GetDBConnection() *sql.DB {
	driverName := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DB_SOURCE")

	Db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("Unable to connect to database")
	}

	return Db
}
