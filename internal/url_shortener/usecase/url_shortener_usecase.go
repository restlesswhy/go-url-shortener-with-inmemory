package usecase

import (
	"context"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
)

type UrlShortenerUC struct {
	shortenerRepo us.UrlShortenerRepository
	cfg *config.Config
}

func NewUrlShortenerUC(cfg *config.Config, shortenerRepo us.UrlShortenerRepository) *UrlShortenerUC {
	return &UrlShortenerUC{
		cfg: cfg,
	}
}

func (u *UrlShortenerUC) Create(ctx context.Context, longUrl string) (string, error) {
	return u.shortenerRepo.Create(ctx, longUrl)
}