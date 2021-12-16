package urlshortener

import "context"

type UrlShortenerUseCase interface {
	Create(ctx context.Context, longUrl string) (string, error)
	Get(ctx context.Context, shortUrl string) (string, error)
}