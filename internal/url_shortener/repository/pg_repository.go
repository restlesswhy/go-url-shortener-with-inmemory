package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
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

func (u *UrlShortenerRepository) CreateRepo(ctx context.Context, longUrl string, shortUrl string) error {
	logger.Info("creating new url in repo")

	query := fmt.Sprintf("INSERT INTO %s (short_url, long_url) VALUES ($1, $2);", urlsTable)
	_, err := u.db.Query(query, shortUrl, longUrl)
	if err != nil {
		return err
	}

	logger.Infof("added: longUrl - %s, shortUrl - %s", longUrl, shortUrl)

	return nil
}


func (u *UrlShortenerRepository) GetRepo(ctx context.Context, longUrl string) (string, error) {
	logger.Info("serching url in repo")

	var resUrl string

	query := fmt.Sprintf("SELECT short_url FROM %s WHERE long_url = $1", urlsTable)
	err := u.db.QueryRow(query, longUrl).Scan(&resUrl)
	if err == sql.ErrNoRows {
		logger.Info("have no url in repo")
		return "", err
	}
	logger.Infof("found this in repo - %s", resUrl)
	return resUrl, err
	
}