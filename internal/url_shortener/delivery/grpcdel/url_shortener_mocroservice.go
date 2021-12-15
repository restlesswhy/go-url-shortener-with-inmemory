package grpcdel

import (
	"context"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
	pb "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
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

	return &pb.UCResponse{
		ShortUrl: "hi there",
	}, nil
}

func (u *UrlShortenerMicroservice) Get(ctx context.Context, in *pb.UGRequest) (*pb.UGResponse, error) {
	return nil, nil
}