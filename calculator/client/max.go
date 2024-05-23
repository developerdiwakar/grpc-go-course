package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})
	// Sending data
	go func() {
		numbers := []int32{4, 7, 3, 5, 12, 13, 34}
		for _, n := range numbers {
			log.Printf("Sending numnber: %d\n", n)
			stream.Send(&pb.MaxRequest{Number: n})
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()
	}()

	// Receving data
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Problem while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum of...: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
