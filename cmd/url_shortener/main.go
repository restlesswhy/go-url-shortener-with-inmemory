package main

import (
	"log"
	"os"

	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/server"
)

func main() {
	log.Println("Starting server")

	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("cant get config: %v", err)
	}

	s := server.NewServer(cfg)
	s.Run()
}