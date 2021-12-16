package urlshortener

import "context"

type UrlShortenerRepository interface {
	CreateRepo(ctx context.Context, longUrl string, shortUrl string) error
	GetRepo(ctx context.Context, longUrl string) (string, error)
}