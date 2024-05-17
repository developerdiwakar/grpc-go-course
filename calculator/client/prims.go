package main

import (
	"context"
	"io"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

func doPrims(c pb.CalculatorServiceClient) {
	log.Printf("doPrims was invoked")
	req := &pb.PrimsRequest{Number: 12390392840}

	stream, err := c.Prims(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Prims: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		log.Printf("Primes: %v", res.Result)
	}
}
