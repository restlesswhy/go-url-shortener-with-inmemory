package main

import (
	"context"
	"log"

	// pb "github.com/restlesswhy/grpc/shortener-client/proto"

	pb "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	client := pb.NewUrlShortenerServiceClient(conn)
	
	resp, err := client.Create(context.Background(), &pb.UCRequest{
		LongUrl: "asdasdasdfdfasd",
	})
	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	// ress, err := client.Get(context.Background(), &pb.UGRequest{
	// 	ShortUrl: "WOEAi4MH23",
	// })
	// if err != nil {
	// 	log.Fatalf("could not get answer: %v", err)
	// }


	log.Println("Short url:", resp.ShortUrl)
	// log.Println("Long url:", ress.LongUrl)
}
