//go:generate mockgen -source pg_repository.go -destination mock/pg_repository.go -package mock
package urlshortener

import (
	"context"

	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
)

type USRepository interface {
	Create(ctx context.Context, longUrl, shortUrl string) error
	Get(ctx context.Context, longUrl, shortUrl string) (models.UrlsLS, bool, error)
}