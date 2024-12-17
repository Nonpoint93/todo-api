package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConfig struct {
	DB *sql.DB
}

func NewDatabase() (*DatabaseConfig, error) {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"done" BOOLEAN
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return &DatabaseConfig{DB: db}, nil
}
