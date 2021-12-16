package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	urlsTable = "urls"
)

type UrlShortenerRepository struct {
	db *sqlx.DB
}

func NewUrlShortenerRepository(db *sqlx.DB) *UrlShortenerRepository {
	return &UrlShortenerRepository{
		db: db,
	}
}

func (u *UrlShortenerRepository) Create(ctx context.Context, longUrl string, shortUrl string) (string, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (short_url, long_url) VALUES ($1, $2) RETURNING id;", urlsTable)
	row := u.db.QueryRow(query, shortUrl, longUrl)

	if err := row.Scan(&id); err != nil {
		return err.Error(), err
	}
	
	fmt.Println(id)
	return urlsTable, nil
}