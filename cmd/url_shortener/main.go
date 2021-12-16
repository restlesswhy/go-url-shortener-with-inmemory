package main

import (
	"log"
	"os"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/server"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/postgres"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/memdb"
)

func main() {
	log.Println("Starting server")

	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("cant get config: %v", err)
	}

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		logger.Fatalf("Postgresql init: %s", err)
	}
	defer psqlDB.Close()

	memdb, err := inmemory.InitMemDB()
	if err != nil {
		logger.Fatalf("Memdb init: %s", err)
	}

	s := server.NewServer(cfg, psqlDB, memdb)
	
	logger.Fatal(s.Run())
}