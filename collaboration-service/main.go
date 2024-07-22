package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"pranjal0207/collaborative-doc-platform/collaboration-service/handler"
	"pranjal0207/collaborative-doc-platform/collaboration-service/model"
	pb "pranjal0207/collaborative-doc-platform/collaboration-service/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	documentStore := model.NewDocumentStore()

	server := &handler.Server{
		DocumentStore: documentStore,
	}

	pb.RegisterCollaborationServiceServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
