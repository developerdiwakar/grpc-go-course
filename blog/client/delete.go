package main

import (
	"context"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog has been invoked----")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Printf("Blog with id: %s was deleted\n", id)
}
