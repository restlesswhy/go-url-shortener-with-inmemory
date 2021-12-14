package server

import (
	"net"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/delivery/grpcdel"
	shortenerService "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"google.golang.org/grpc"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	
	server := grpc.NewServer()
	shortener := grpcdel.NewUrlShortenerMicroservice(s.cfg)
	shortenerService.RegisterUrlShortenerServiceServer(server, shortener)
	
	server.Serve(l)
	return nil
}