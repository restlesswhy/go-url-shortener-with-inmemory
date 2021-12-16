package usecase

import (
	"context"
	"database/sql"

	"github.com/cristalhq/base64"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
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
	
	shortUrl, err := u.memdb.GetLongInmemory(longUrl)
	if err != nil {
		return shortUrl, err
	}

	if shortUrl == "" {
		shortUrl, err = u.shortenerRepo.GetRepo(ctx, longUrl)
		if err == nil {
			if err := u.memdb.CreateInmemory(shortUrl, longUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
		}

		if err == sql.ErrNoRows {
			shortUrl = getUniqueString(longUrl)
			
			if err := u.memdb.CreateInmemory(shortUrl, longUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
			if err := u.shortenerRepo.CreateRepo(ctx, longUrl, shortUrl); err != nil {
				// u.memdb.Printmem(shortUrl+"aasf3")
				return shortUrl, errors.Wrap(err, "u.shortenerRepo.CreateRepo")
			}
			
		}
	}
	return shortUrl, nil
}

func (u *UrlShortenerUC) Get(ctx context.Context, shortUrl string) (string, error) {
	// longUrl := u.memdb.GetInmemory()
	return "", nil
}

func getUniqueString(longUrl string) string {
	shortUrl := base64.RawURLEncoding.EncodeStringToString(longUrl)
	return shortUrl[len(shortUrl)-10:]
}

