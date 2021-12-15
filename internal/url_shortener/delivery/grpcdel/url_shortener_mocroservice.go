package grpcdel

import (
	"context"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	pb "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
)

type UrlShortenerMicroservice struct {
	pb.UnimplementedUrlShortenerServiceServer
	cfg *config.Config
	
}

func NewUrlShortenerMicroservice(cfg *config.Config) *UrlShortenerMicroservice{
	return &UrlShortenerMicroservice{
		cfg: cfg,
	}
}

func (u *UrlShortenerMicroservice) Create(ctx context.Context, in *pb.UCRequest) (*pb.UCResponse, error) {
	return nil, nil
}

func (u *UrlShortenerMicroservice) Get(ctx context.Context, in *pb.UGRequest) (*pb.UGResponse, error) {
	return nil, nil
}