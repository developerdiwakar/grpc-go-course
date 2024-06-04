package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/developerdiwakar/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("---DeleteBlog was invoked---")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot parse id %v", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete blog with given id %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot was not found",
		)
	}

	return &emptypb.Empty{}, nil
}
