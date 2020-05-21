package database

import "database/sql"

func GetConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=Joe dbname=test sslmode=disable port=26257")

	return db, err
}
