package main

import (
	"log"
	"net"

	pb "github.com/developerdiwakar/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Println("Listening on port:", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	// gRPC reflection & Evans CLI
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalln("Failed to server: ", err)
	}

}
