package main

import (
	"context"
	"log"
	"time"

	pb "github.com/developerdiwakar/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet function invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Diwakar"},
		{FirstName: "Prabhakar"},
		{FirstName: "Viraj"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
