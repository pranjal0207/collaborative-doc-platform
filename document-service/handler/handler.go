package handler

import (
	"context"

	"pranjal0207/collaborative-doc-platform/document-service/model"
	pb "pranjal0207/collaborative-doc-platform/document-service/proto"

	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedDocumentServiceServer
	DocumentModel *model.DocumentModel
}

func (s *Server) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.CreateDocumentResponse, error) {
	documentID := uuid.New().String()

	document := model.Document{
		Title:      req.Title,
		Author:     req.Author,
		Content:    "",
		DocumentID: documentID,
	}

	err := s.DocumentModel.CreateDocument(ctx, &document)

	if err != nil {
		return nil, err
	}

	return &pb.CreateDocumentResponse{
		DocumentId: documentID,
	}, nil
}

func (s *Server) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	document, err := s.DocumentModel.GetDocumentByID(ctx, req.DocumentId)

	if err != nil {
		return nil, err
	}

	return &pb.GetDocumentResponse{
		DocumentId: document.DocumentID,
		Title:      document.Title,
		Content:    document.Content,
		Author:     document.Author,
		Versions:   document.Versions,
	}, nil
}

func (s *Server) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	err := s.DocumentModel.DeleteDocumentByID(ctx, req.DocumentId)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteDocumentResponse{
		DocumentId: req.DocumentId,
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
