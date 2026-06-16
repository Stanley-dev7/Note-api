package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite", "./notes.db")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not reachable:", err)
	}

	log.Println("Database connected successfully")
}

// CREATE TABLES
func InitTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	noteTable := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal("User table error:", err)
	}

	_, err = DB.Exec(noteTable)
	if err != nil {
		log.Fatal("Notes table error:", err)
	}

	log.Println("Tables created successfully")
}