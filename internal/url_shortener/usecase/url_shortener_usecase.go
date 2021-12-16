package usecase

import (
	"context"
	"database/sql"

	"github.com/cristalhq/base64"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
)

type UrlShortenerUC struct {
	shortenerRepo us.UrlShortenerRepository
	cfg *config.Config
	memdb us.UrlShortenerInmemory
}

func NewUrlShortenerUC(cfg *config.Config, shortenerRepo us.UrlShortenerRepository, memdb us.UrlShortenerInmemory) *UrlShortenerUC {
	return &UrlShortenerUC{
		shortenerRepo: shortenerRepo,
		cfg: cfg,
		memdb: memdb,
	}
}

func (u *UrlShortenerUC) Create(ctx context.Context, longUrl string) (string, error) {
	

	shortUrl, err := u.memdb.GetShortInmemory(longUrl)
	if err != nil {
		return shortUrl, err
	}

	if shortUrl == "" {
		shortUrl = getUniqueString(longUrl)
		urls, err := u.shortenerRepo.GetRepo(ctx, longUrl, shortUrl)
		if err == nil {
			if err := u.memdb.CreateInmemory(urls.ShortUrl, urls.LongUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
		}
		logger.Infof("urls in model: %s, %s", urls.LongUrl, urls.ShortUrl)
		if err == sql.ErrNoRows {
			if err := u.memdb.CreateInmemory(shortUrl, longUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
			if err := u.shortenerRepo.CreateRepo(ctx, longUrl, shortUrl); err != nil {
				return shortUrl, errors.Wrap(err, "u.shortenerRepo.CreateRepo")
			}
			
		}
	}
	return shortUrl, nil
}

func (u *UrlShortenerUC) Get(ctx context.Context, shortUrl string) (string, error) {
	longUrl, err := u.memdb.GetLongInmemory(shortUrl)
	if err != nil {
		return shortUrl, err
	}

	if longUrl == "" {
		urls, err := u.shortenerRepo.GetRepo(ctx, longUrl, shortUrl)
		if err == nil {
			if err := u.memdb.CreateInmemory(urls.ShortUrl, urls.LongUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
		}
		if err == sql.ErrNoRows {
			return "have no long url", nil
		}
		longUrl = urls.LongUrl

	}

	return longUrl, nil
}

func getUniqueString(longUrl string) string {
	shortUrl := base64.RawURLEncoding.EncodeStringToString(longUrl)
	return shortUrl[len(shortUrl)-10:]
}

