package database

import (
	"database/sql"
	"log"
	"time"
)

type DomainDB struct {
	Host             string
	SslGrade         string
	PreviousSslGrade string
	LastSearch       time.Time
}

type domainRepo struct {
	db *sql.DB
}

// NewRepository creates a crockoach repository with the necessary dependencies
func NewDomainRepository(db *sql.DB) domainRepo {
	return domainRepo{db: db}
}

func (r domainRepo) CreateDomain(d *DomainDB) error {
	sqlQuery := `INSERT INTO androidqueries.domain (host, sslgrade, sslpreviousgrade, lastsearch) 
	VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(sqlQuery, d.Host, d.SslGrade, d.PreviousSslGrade)
	if err != nil {
		return err
	}
	return nil
}

func (r domainRepo) FetchDomain(hostname string) (*DomainDB, error) {
	sqlQuery := `SELECT host, sslgrade, sslpreviousgrade, lastsearch FROM androidqueries.domain WHERE (host) VALUES ($1)`
	_, err := r.db.Exec(sqlQuery, hostname)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var host, sslgrade, previoussslgrade string
	var lastsearch time.Time

	for rows.Next() {
		if err := rows.Scan(&host, &sslgrade, &previoussslgrade, &lastsearch); err != nil {
			log.Fatal(err)
		}

		return &DomainDB{host, sslgrade, previoussslgrade, lastsearch}, nil
	}

	return nil, err
}

func (r domainRepo) UpdateDomain(d *DomainDB) error {
	sqlQuery := `UPDATE androidqueries.domain SET sslgrade =  $1, sslpreviousgrade = $2 , lastsearch = NOW() WHERE host = $3
	VALUES ($1, $2, $3, NOW()`
	_, err := r.db.Exec(sqlQuery, d.SslGrade, d.PreviousSslGrade, d.Host)
	if err != nil {
		return err
	}

	return nil
}
