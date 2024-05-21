package main

import (
	"context"
	"log"
	"time"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

func getAvg(c pb.CalculatorServiceClient) {
	log.Println("Function getAvg invoked.")
	reqs := []int32{1, 2, 3, 4}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling getAvg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending requst: %v\n", req)
		stream.Send(&pb.AvgRequest{Number: req})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Error while receiving response from server: %v\n", err)
	}

	log.Printf("Avg: %v", res.Result)
}
