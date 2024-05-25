package main

import (
	"context"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog invoked")

	blog := &pb.Blog{
		AuthorId: "diwakar",
		Title:    "my first blog",
		Content:  "Content of the first blog post",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog Created: %s\n", res.Id)

	return res.Id
}
