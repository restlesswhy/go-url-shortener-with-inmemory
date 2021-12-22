package grpcdel

import (
	"context"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
	pb "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
)

type USMicroservice struct {
	pb.UnimplementedUrlShortenerServiceServer
	cfg *config.Config
	shortenerUC us.USUseCase
}

func NewUSMicroservice(cfg *config.Config, shortenerUC us.USUseCase) *USMicroservice{
	return &USMicroservice{
		cfg: cfg,
		shortenerUC: shortenerUC,
	}
}

func (u *USMicroservice) Create(ctx context.Context, in *pb.UCRequest) (*pb.UCResponse, error) {
	shortUrl, err := u.shortenerUC.Create(ctx, in.LongUrl)
	if err != nil {
		logger.Fatal(err)
	}

	return &pb.UCResponse{
		ShortUrl: shortUrl,
	}, nil
}

func (u *USMicroservice) Get(ctx context.Context, in *pb.UGRequest) (*pb.UGResponse, error) {
	longUrl, err := u.shortenerUC.Get(ctx, in.ShortUrl)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return &pb.UGResponse{
		LongUrl: longUrl,
	}, nil
}