package urlshortener

import "context"

type UrlShortenerUseCase interface {
	Create(ctx context.Context, longUrl string) (string, error)
}