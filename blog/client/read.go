package main

import (
	"context"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("---readBlog was invoked---")
	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading blog: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
