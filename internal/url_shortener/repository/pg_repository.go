package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
)

const (
	urlsTable = "urls"
)

type USRepository struct {
	db *sqlx.DB
}

func NewUSRepository(db *sqlx.DB) *USRepository {
	return &USRepository{
		db: db,
	}
}

// CreateRepo insert url in repository
func (u *USRepository) Create(ctx context.Context, longUrl string, shortUrl string) error {
	logger.Info("Creating new url in repo")

	query := fmt.Sprintf("INSERT INTO %s (short_url, long_url) VALUES ($1, $2);", urlsTable)
	_, err := u.db.Exec(query, shortUrl, longUrl)
	if err != nil {
		return errors.Wrap(err, "Create.Exec")
	}

	logger.Infof("Created in repo: longUrl - %s, shortUrl - %s", longUrl, shortUrl)
	return nil
}

// GetRepo search url in repository
func (u *USRepository) Get(ctx context.Context, longUrl, shortUrl string) (models.UrlsLS, bool, error) {
	logger.Info("Serching url in repo...")

	var isExist bool
	var urls models.UrlsLS

	query := fmt.Sprintf("SELECT short_url, long_url FROM %s WHERE long_url = $1 OR short_url = $2", urlsTable)
	err := u.db.Get(&urls, query, longUrl, shortUrl)
	if err == sql.ErrNoRows {
		logger.Info("Have no url in repo")
		return urls, isExist, nil
	}
	if err != nil {
		return urls, isExist, errors.Wrap(err, "Get.Get")
	}

	isExist = true

	logger.Infof("Found this in repo")
	return urls, isExist, nil
}