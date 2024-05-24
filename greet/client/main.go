package main

import (
	"log"
	"time"

	pb "github.com/developerdiwakar/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	// Dialing TCP call without SSL
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
	doGreetWithDeadline(c, 2*time.Second) // Because durartion < 3, Should output: Deadline Exceeded!
}
