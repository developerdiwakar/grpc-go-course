package main

import (
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

// Prim's Algorithm (pseudo code):
// k = 2
// N = 210
// while N > 1:
//     if N % k == 0:   // if k evenly divides into N
//         print k      // this is a factor
//         N = N / k    // divide N by k so that we have the rest of the number left.
//     else:
//         k = k + 1

func (s *Server) Prims(in *pb.PrimsRequest, stream pb.CalculatorService_PrimsServer) error {
	log.Printf("Prims Function was invoked with: %v\n", in)

	num := in.Number
	divisor := int64(2)

	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&pb.PrimsResponse{Result: divisor})
			num = num / divisor
		} else {
			divisor++
		}

	}

	return nil

}
