package main

import (
	"context"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum function was invoked with: ", in)
	sum := in.FirstNumber + in.SecondNumber
	return &pb.SumResponse{
		Result: sum,
	}, nil
}
