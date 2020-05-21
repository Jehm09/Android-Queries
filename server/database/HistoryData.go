package database

import (
	"database/sql"
	"log"
)

type HistoryDB struct {
	Items []string
}

type historyRepo struct {
	db *sql.DB
}

// NewRepository creates a crockoach repository with the necessary dependencies
func NewHistoyRepository(db *sql.DB) historyRepo {
	return historyRepo{db: db}
}

func (r domainRepo) CreateHistory(hostname string) error {
	sqlQuery := `INSERT INTO androidqueries.history (host) 
	VALUES ($1)`
	_, err := r.db.Exec(sqlQuery, hostname)
	if err != nil {
		return err
	}
	return nil
}

func (r domainRepo) FetchHistory() (*HistoryDB, error) {
	sqlQuery := `SELECT host FROM androidqueries.history`

	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []string

	for rows.Next() {
		var host string
		if err := rows.Scan(&host); err != nil {
			log.Fatal(err)
		}

		items = append(items, host)
	}

	return &HistoryDB{Items: items}, nil
}
