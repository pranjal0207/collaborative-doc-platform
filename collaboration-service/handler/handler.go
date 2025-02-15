package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"pranjal0207/collaborative-doc-platform/collaboration-service/model"
	pb "pranjal0207/collaborative-doc-platform/collaboration-service/proto"
	"pranjal0207/collaborative-doc-platform/collaboration-service/utils"

	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedCollaborationServiceServer
	DocumentStore *model.DocumentStore
}

func (s *Server) JoinDocument(ctx context.Context, req *pb.JoinDocumentRequest) (*pb.JoinDocumentResponse, error) {
	doc, err := s.DocumentStore.GetDocument(req.DocumentId)
	if (err != nil) {
		return nil, err
	}

	sessionID := uuid.New().String()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := utils.Upgrade(w, r)
		if err != nil {
			http.Error(w, "Failed to upgrade to websocket", http.StatusInternalServerError)
			return
		}
		user := &model.User{
			UserID:     req.UserId,
			Connection: conn,
		}
		doc.Mutex.Lock()
		doc.Users[req.UserId] = user
		doc.Mutex.Unlock()

		go s.handleUserConnection(doc, user)
	})

	log.Printf("User %s joined document %s", req.UserId, req.DocumentId)
	return &pb.JoinDocumentResponse{SessionId: sessionID}, nil
}

func (s *Server) SyncChanges(ctx context.Context, req *pb.SyncChangesRequest) (*pb.SyncChangesResponse, error) {
	doc, err := s.DocumentStore.GetDocument(req.DocumentId)
	if (err != nil) {
		return nil, err
	}

	var changes map[string]interface{}
	if err := json.Unmarshal([]byte(req.Changes), &changes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal changes: %v", err)
	}

	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	doc.Content = req.Changes
	go s.DocumentStore.UpdateDocument(req.DocumentId, req.Changes)

	for _, user := range doc.Users {
		if user.UserID != req.UserId {
			if err := user.Connection.WriteMessage(websocket.TextMessage, []byte(req.Changes)); err != nil {
				log.Printf("Failed to send message to user %s: %v", user.UserID, err)
			}
		}
	}

	return &pb.SyncChangesResponse{Success: true}, nil
}

func (s *Server) LeaveDocument(ctx context.Context, req *pb.LeaveDocumentRequest) (*pb.LeaveDocumentResponse, error) {
	doc, err := s.DocumentStore.GetDocument(req.DocumentId)
	if (err != nil) {
		return nil, err
	}

	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	if user, exists := doc.Users[req.UserId]; exists {
		user.Connection.Close()
		delete(doc.Users, req.UserId)
		log.Printf("User %s left document %s", req.UserId, req.DocumentId)
	}

	return &pb.LeaveDocumentResponse{Success: true}, nil
}

func (s *Server) handleUserConnection(doc *model.Document, user *model.User) {
	defer func() {
		doc.Mutex.Lock()
		delete(doc.Users, user.UserID)
		doc.Mutex.Unlock()
		user.Connection.Close()
		log.Printf("Connection closed for user %s", user.UserID)
	}()

	for {
		_, message, err := user.Connection.ReadMessage()
		if err != nil {
			log.Printf("Read error for user %s: %v", user.UserID, err)
			break
		}
		log.Printf("Received message from user %s: %s", user.UserID, message)
	}
}
