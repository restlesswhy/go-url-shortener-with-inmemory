package usecase

import (
	"context"
	"fmt"
	// "fmt"

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

	shortUrl := "asdads"
	fmt.Println(shortUrl)
	x, err := u.shortenerRepo.Create(ctx, longUrl, shortUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
	// return u.shortenerRepo.Create(ctx, longUrl, shortUrl)
	return x, nil
}