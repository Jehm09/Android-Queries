package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=test dbname=androidqueries sslmode=disable port=26257")

	return db, err
}
