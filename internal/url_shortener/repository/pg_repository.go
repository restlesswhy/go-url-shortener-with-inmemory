package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
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

// CreateRepo insert url in repository
func (u *UrlShortenerRepository) CreateRepo(ctx context.Context, longUrl string, shortUrl string) error {
	logger.Info("creating new url in repo")

	query := fmt.Sprintf("INSERT INTO %s (short_url, long_url) VALUES ($1, $2);", urlsTable)
	_, err := u.db.Query(query, shortUrl, longUrl)
	if err != nil {
		return err
	}

	logger.Infof("added to repo: longUrl - %s, shortUrl - %s", longUrl, shortUrl)

	return nil
}

// GetRepo search url in repository
func (u *UrlShortenerRepository) GetRepo(ctx context.Context, longUrl, shortUrl string) (models.UrlsLS, bool) {
	logger.Info("serching url in repo")
	var isExist bool
	var urls models.UrlsLS
	// var resUrl string

	query := fmt.Sprintf("SELECT short_url, long_url FROM %s WHERE long_url = $1 OR short_url = $2", urlsTable)
	err := u.db.Get(&urls, query, longUrl, shortUrl)

	// err := u.db.QueryRow(query, longUrl, shortUrl).Scan(&resUrl)
	if err == sql.ErrNoRows {
		logger.Info("have no url in repo")
		return urls, isExist
	}
	if err != nil {
		logger.Errorf("select error: %s", err)
		return urls, isExist
	}

	isExist = true

	logger.Infof("found this in repo - %s", urls.ShortUrl)
	return urls, isExist
	
}