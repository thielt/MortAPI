package mortapi

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./sql/users.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the database connection is valid by pinging it
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Create Users table if none found
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}
