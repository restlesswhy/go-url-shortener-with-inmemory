package usecase

import (
	"context"

	"github.com/cristalhq/base64"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
)

type UrlShortenerUC struct {
	shortenerRepo us.UrlShortenerRepository
	cfg *config.Config
}

func NewUrlShortenerUC(cfg *config.Config, shortenerRepo us.UrlShortenerRepository) *UrlShortenerUC {
	return &UrlShortenerUC{
		shortenerRepo: shortenerRepo,
		cfg: cfg,
	}
}

func (u *UrlShortenerUC) Create(ctx context.Context, longUrl string) (string, error) {

	shortUrl := getUniqueString(longUrl)

	urlIsExist, err := u.shortenerRepo.Create(ctx, longUrl, shortUrl)
	if urlIsExist && err == nil {
		// TODO добавляем в in-memory
		return shortUrl, err
	}

	if err != nil {
		return shortUrl, errors.Wrap(err, "u.shortenerRepo.Create")
	}
	
	return shortUrl, nil
}

func getUniqueString(longUrl string) string {
	shortUrl := base64.RawURLEncoding.EncodeStringToString(longUrl)
	return shortUrl[len(shortUrl)-10:]
}

