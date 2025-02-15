package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"pranjal0207/collaborative-doc-platform/user-service/handler"
	"pranjal0207/collaborative-doc-platform/user-service/model"
	pb "pranjal0207/collaborative-doc-platform/user-service/proto"
	"pranjal0207/collaborative-doc-platform/user-service/utils"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dynamodb, err := utils.DynamoDBInstance()

	if err != nil {
		log.Fatalf("failed to establish connection to db: %v", err)
	}

	pb.RegisterUserServiceServer(s, &handler.Server{UserModel: &model.UserModel{
		DynamoDB:  dynamodb,
		TableName: "doc_users",
	}})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
