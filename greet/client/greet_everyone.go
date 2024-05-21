package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/developerdiwakar/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone function invoked.")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Printf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Radhika"}, {FirstName: "Sonali"}, {FirstName: "Marie"},
	}

	waitc := make(chan int)

	go func() {
		for _, req := range reqs {
			log.Printf("Sending Request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving data from the server: %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
