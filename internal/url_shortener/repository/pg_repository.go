package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (u *UrlShortenerRepository) Create(ctx context.Context, longUrl string, shortUrl string) (bool, error) {
	var resUrl string
	var urlIsExist bool

	query := fmt.Sprintf("SELECT short_url FROM %s WHERE short_url = $1 OR long_url = $2", urlsTable)
	err := u.db.QueryRow(query, shortUrl, longUrl).Scan(&resUrl)
	if err == nil {
		urlIsExist = true
	}

	if err == sql.ErrNoRows {
		query := fmt.Sprintf("INSERT INTO %s (short_url, long_url) VALUES ($1, $2);", urlsTable)
		_, err := u.db.Query(query, shortUrl, longUrl)
		return urlIsExist, err
	} else if err != nil {
		return urlIsExist, err
	}
	
	return urlIsExist, errors.New("this url is already create")
}