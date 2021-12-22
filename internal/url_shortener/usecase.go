//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock
package urlshortener

import "context"

type UrlShortenerUseCase interface {
	Create(ctx context.Context, longUrl string) (string, error)
	Get(ctx context.Context, shortUrl string) (string, error)
}