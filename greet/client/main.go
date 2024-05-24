package main

import (
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {

	tls := false // Change it to false/true if needed
	opts := []grpc.DialOption{}
	//TODO: Could not greet: rpc error: code = Unavailable desc = connection error:
	// desc = "transport: authentication handshake failed:
	// tls: failed to verify certificate: x509: cannot validate certificate for 0.0.0.0 because it doesn't contain any IP SANs"
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 2*time.Second) // Because durartion < 3, Should output: Deadline Exceeded!
}
