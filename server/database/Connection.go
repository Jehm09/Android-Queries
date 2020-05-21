package database

import "database/sql"

func GetConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=Joe dbname=androidqueries sslmode=disable port=26257")

	return db, err
}
