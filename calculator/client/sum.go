package main

import (
	"context"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 3, SecondNumber: 10})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum = %v\n", res.Result)
}
