package main

import (
	"context"
	"log"

	shortenerService "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := shortenerService.NewUrlShortenerServiceClient(conn)


	resp, err := client.Create(context.Background(), &shortenerService.UCRequest{
		LongUrl: "asdasdasdasdasd",
	})
	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	log.Println("Short url:", resp.ShortUrl)
}