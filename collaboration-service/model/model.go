package model

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Document struct {
	DocumentID string
	Content    string
	Users      map[string]*User // Map of user_id to User
	Mutex      sync.Mutex
}

type User struct {
	UserID     string
	Connection *websocket.Conn
}

type DocumentStore struct {
	Documents map[string]*Document
	Mutex     sync.Mutex
}

func NewDocumentStore() *DocumentStore {
	return &DocumentStore{
		Documents: make(map[string]*Document),
	}
}

func (store *DocumentStore) GetDocument(documentID string) *Document {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()
	if doc, exists := store.Documents[documentID]; exists {
		return doc
	}
	doc := &Document{
		DocumentID: documentID,
		Users:      make(map[string]*User),
	}
	store.Documents[documentID] = doc
	return doc
}
