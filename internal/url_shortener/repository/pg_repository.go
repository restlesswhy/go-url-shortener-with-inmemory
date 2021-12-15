package repository

import "github.com/jmoiron/sqlx"

type UrlShortenerRepository struct {
	db *sqlx.DB
}

func NewUrlShortenerRepository(db *sqlx.DB) *UrlShortenerRepository {
	return &UrlShortenerRepository{
		db: db,
	}
}

