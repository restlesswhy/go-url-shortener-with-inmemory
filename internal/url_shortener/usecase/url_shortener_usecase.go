package usecase

import "github.com/restlesswhy/grpc/url-shortener-microservice/config"

type UrlShortenerUC struct {
	cfg *config.Config
}

func NewUrlShortenerUC(cfg *config.Config) *UrlShortenerUC {
	return &UrlShortenerUC{
		cfg: cfg,
	}
}