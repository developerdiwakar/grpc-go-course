package main

import (
	"context"
	"io"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog has been invoked---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something went wrong! %v\n", err)
		}

		log.Println(res)
	}
}
