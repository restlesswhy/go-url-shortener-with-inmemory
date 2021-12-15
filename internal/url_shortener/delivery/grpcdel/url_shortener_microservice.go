package grpcdel

import (
	"context"
	"fmt"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
	pb "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"google.golang.org/grpc/status"
)

type UrlShortenerMicroservice struct {
	pb.UnimplementedUrlShortenerServiceServer
	cfg *config.Config
	shortenerUC us.UrlShortenerUseCase
}

func NewUrlShortenerMicroservice(cfg *config.Config, shortenerUC us.UrlShortenerUseCase) *UrlShortenerMicroservice{
	return &UrlShortenerMicroservice{
		cfg: cfg,
		shortenerUC: shortenerUC,
	}
}

func (u *UrlShortenerMicroservice) Create(ctx context.Context, in *pb.UCRequest) (*pb.UCResponse, error) {
	shortUrl, err := u.shortenerUC.Create(ctx, in.LongUrl)
	if err != nil {
		return nil, status.Errorf(400, "u.shortenerUC.Create: %v", err)	
	}

	return &pb.UCResponse{
		ShortUrl: fmt.Sprintf("Short url: %s", shortUrl),
	}, nil
}

func (u *UrlShortenerMicroservice) Get(ctx context.Context, in *pb.UGRequest) (*pb.UGResponse, error) {
	return nil, nil
}