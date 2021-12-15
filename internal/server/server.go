package server

import (
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/delivery/grpcdel"
	shortenerService "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/repository"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/usecase"
	"google.golang.org/grpc"
)

type Server struct {
	cfg *config.Config
	db *sqlx.DB
}

func NewServer(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{
		cfg: cfg,
		db: db,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer l.Close()
	
	shortenerRepository := repository.NewUrlShortenerRepository(s.db)
	shortenerUseCase := usecase.NewUrlShortenerUC(s.cfg, shortenerRepository)

	server := grpc.NewServer()
	shortener := grpcdel.NewUrlShortenerMicroservice(s.cfg)
	shortenerService.RegisterUrlShortenerServiceServer(server, shortener)
	
	server.Serve(l)
	return nil
}