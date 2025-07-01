package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func SetupDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

	fmt.Println("Connected to SQLite database!")
	return db
}
