package urlshortener

import (
	"context"

	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
)

type UrlShortenerRepository interface {
	CreateRepo(ctx context.Context, longUrl, shortUrl string) error
	GetRepo(ctx context.Context, longUrl, shortUrl string) (models.UrlsLS, error)
}