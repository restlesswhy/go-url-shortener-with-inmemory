package main

import (
	"log"
	"os"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/server"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/postgres"
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

	s := server.NewServer(cfg, psqlDB)
	
	logger.Fatal(s.Run())
}