package main

import (
	"io"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

// Example:

// The client will send a stream of number (1,2,3,4) and the server will respond with (2.5),
// because (1+2+3+4)/4 = 2.5

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Function Avg was invoked")

	var res float64
	var sum int32
	len := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res = float64(sum) / float64(len)
			return stream.SendAndClose(&pb.AvgResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receving the number: %d\n", req.Number)
		len++
		sum += req.Number
	}

}
