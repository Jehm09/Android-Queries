package database

import "time"

type DomainDB struct {
	Host             string
	SslGrade         string
	SslPreviousGrade string
	LastSearch       time.Time
}
