package main

import (
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50052"

func main() {
	// Dialing TCP call without SSL
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrims(c)
	// getAvg(c)
	// doMax(c)
	doSqrt(c, 3)
	doSqrt(c, -3) // invalid argument
}
