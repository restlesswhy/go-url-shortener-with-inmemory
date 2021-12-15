package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UrlShortenerRepository struct {
	db *sqlx.DB
}

func NewUrlShortenerRepository(db *sqlx.DB) *UrlShortenerRepository {
	return &UrlShortenerRepository{
		db: db,
	}
}

func (u *UrlShortenerRepository) Create(ctx context.Context, longUrl string) (string, error) {
	
	return "", nil
}