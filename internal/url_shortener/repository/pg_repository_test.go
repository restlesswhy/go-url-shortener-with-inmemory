package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
	"github.com/stretchr/testify/require"
)


func TestUrlShortenerRepository_CreateRepo(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	shortRepo := NewUrlShortenerRepository(sqlxDB)

	// columns := []string{"short_url", "long_url"}
	mockUrl := models.UrlsLS{
		ShortUrl: "as;dasd;asld",
		LongUrl: "https://habr.com/ru/company/vk/blog/476276/",
	}

	// rows := sqlmock.NewRows(columns).AddRow(
	// 	mockUrl.ShortUrl,
	// 	mockUrl.LongUrl,
	// )
	query := "INSERT INTO urls (short_url, long_url) VALUES ($1, $2);"
	// prep := mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(mockUrl.ShortUrl, mockUrl.LongUrl).WillReturnResult(sqlmock.NewResult(0, 0))

	err = shortRepo.CreateRepo(context.Background(), mockUrl.LongUrl, mockUrl.ShortUrl)

	require.NoError(t, err)
}

func TestUrlShortenerRepository_GetRepo(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	shortRepo := NewUrlShortenerRepository(sqlxDB)

	columns := []string{"short_url", "long_url"}
	mockUrl := models.UrlsLS{
		ShortUrl: "as;dasd;asld",
		LongUrl: "https://habr.com/ru/company/vk/blog/476276/",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUrl.ShortUrl,
		mockUrl.LongUrl,
	)
	query := "SELECT short_url, long_url FROM urls WHERE long_url = $1 OR short_url = $2"
	// prep := mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(mockUrl.LongUrl, mockUrl.ShortUrl).WillReturnRows(rows)

	urls, isExist := shortRepo.GetRepo(context.Background(), mockUrl.LongUrl, mockUrl.ShortUrl)

	require.NotNil(t, urls)
	require.Equal(t, mockUrl.ShortUrl, urls.ShortUrl)
	require.True(t, isExist)
}

func TestUrlShortenerRepository_GetRepoNil(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	shortRepo := NewUrlShortenerRepository(sqlxDB)

	// columns := []string{"short_url", "long_url"}
	mockUrl := models.UrlsLS{
		ShortUrl: "as;dasd;asld",
		LongUrl: "https://habr.com/ru/company/vk/blog/476276/",
	}

	// rows := sqlmock.NewRows(columns).AddRow(
	// 	mockUrl.ShortUrl,
	// 	mockUrl.LongUrl,
	// )
	query := "SELECT short_url, long_url FROM urls WHERE long_url = $1 OR short_url = $2"
	// prep := mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(mockUrl.LongUrl, mockUrl.ShortUrl).WillReturnError(sql.ErrNoRows)

	urls, isExist := shortRepo.GetRepo(context.Background(), mockUrl.LongUrl, mockUrl.ShortUrl)

	require.Equal(t, models.UrlsLS{}, urls)
	// require.Equal(t, mockUrl.ShortUrl, urls.ShortUrl)
	require.False(t, isExist)
}