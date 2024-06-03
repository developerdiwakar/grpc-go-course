package main

import (
	"context"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---UpdateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Diwakar",
		Title:    "A new title",
		Content:  "Content for blog with some awesome addition.",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating the blog: %v\n", err)
	}

	log.Println("Blog has been updated")
}
