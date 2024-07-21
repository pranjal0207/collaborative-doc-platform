package handler

import (
	"context"
	"pranjal0207/collaborative-doc-platform/document-service/model"
	pb "pranjal0207/collaborative-doc-platform/document-service/proto"
)

type Server struct {
	pb.UnimplementedDocumentServiceServer
	DocumentModel *model.DocumentModel
}

func (s *Server) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.CreateDocumentResponse, error) {
	return &pb.CreateDocumentResponse{
		DocumentId: "SOME ID",
	}, nil
}

func (s *Server) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	return &pb.GetDocumentResponse{
		DocumentId: "SOME ID",
		Title: "Title",
    	Content:"Content",
    	Author:"Authot",
    	Versions: []string {"version"},
	}, nil
}

func (s *Server) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	return &pb.DeleteDocumentResponse{
		DocumentId: "SOME ID",
	}, nil
}

func (s *Server) UpdateDocument(ctx context.Context, req *pb.UpdateDocumentRequest) (*pb.UpdateDocumentResponse, error) {
	return &pb.UpdateDocumentResponse{
		DocumentId: "SOME ID",
	}, nil
}

func (s *Server) ListDocumentVersions(ctx context.Context, req *pb.ListDocumentVersionsRequest) (*pb.ListDocumentVersionsResponse, error) {
	return &pb.ListDocumentVersionsResponse{
		Versions: []string{"version1", "version2"},
	}, nil
}

