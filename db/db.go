package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	var err error

	dsn := os.Getenv("DATABASE_URL")

	DB, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to connect to PostgresSQL:", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY
	title TEXT NOT NULL
	author TEXT NOT NULL
	year INT NOT NULL
	);
	`

	DB.MustExec(schema)
}
