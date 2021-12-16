package urlshortener

import "context"

type UrlShortenerRepository interface {
	Create(ctx context.Context, longUrl string, shortUrl string) (bool, error)
}