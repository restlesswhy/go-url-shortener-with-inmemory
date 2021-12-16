package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/go-memdb"
	"github.com/jmoiron/sqlx"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/delivery/grpcdel"
	shortenerService "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/repository"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/inmemory"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/usecase"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	cfg *config.Config
	db *sqlx.DB
	memdb *memdb.MemDB
}

func NewServer(cfg *config.Config, db *sqlx.DB, memdb *memdb.MemDB) *Server {
	return &Server{
		cfg: cfg,
		db: db,
		memdb: memdb,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer l.Close()
	
	shortenerInmemory := inmemory.NewUrlShortenerInmemory(s.memdb)
	shortenerRepository := repository.NewUrlShortenerRepository(s.db)
	shortenerUseCase := usecase.NewUrlShortenerUC(s.cfg, shortenerRepository, shortenerInmemory)

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.Server.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
		Time:              s.cfg.Server.Timeout * time.Minute,
	}))
	shortener := grpcdel.NewUrlShortenerMicroservice(s.cfg, shortenerUseCase)
	shortenerService.RegisterUrlShortenerServiceServer(server, shortener)
	
	go func() {
		logger.Infof("Server is listening on port: %v", s.cfg.Server.Port)
		logger.Fatal(server.Serve(l))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	<-quit
	

	server.GracefulStop()
	logger.Info("Server Exited Properly")

	return nil
}