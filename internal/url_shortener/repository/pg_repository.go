package repository

import "github.com/jmoiron/sqlx"

type UrlShortenerRepository struct {
	db *sqlx.DB
}