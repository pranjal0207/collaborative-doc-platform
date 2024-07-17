package main

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

const (
    userServiceEndpoint         = "localhost:50051"
    documentServiceEndpoint     = "localhost:50052"
    collaborationServiceEndpoint = "localhost:50053"
    chatServiceEndpoint         = "localhost:50054"
    notificationServiceEndpoint = "localhost:50055"
    analyticsServiceEndpoint    = "localhost:50056"
    searchServiceEndpoint       = "localhost:50057"
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()

    opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

    // Register User Service
    if err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userServiceEndpoint, opts); err != nil {
        return err
    }
    // Register Document Service
    if err := document.RegisterDocumentServiceHandlerFromEndpoint(ctx, mux, documentServiceEndpoint, opts); err != nil {
        return err
    }
    // Register Collaboration Service
    if err := collaboration.RegisterCollaborationServiceHandlerFromEndpoint(ctx, mux, collaborationServiceEndpoint, opts); err != nil {
        return err
    }
    // Register Chat Service
    if err := chat.RegisterChatServiceHandlerFromEndpoint(ctx, mux, chatServiceEndpoint, opts); err != nil {
        return err
    }
    // Register Notification Service
    if err := notification.RegisterNotificationServiceHandlerFromEndpoint(ctx, mux, notificationServiceEndpoint, opts); err != nil {
        return err
    }
    // Register Analytics Service
    if err := analytics.RegisterAnalyticsServiceHandlerFromEndpoint(ctx, mux, analyticsServiceEndpoint, opts); err != nil {
        return err
    }
    // Register Search Service
    if err := search.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, searchServiceEndpoint, opts); err != nil {
        return err
    }

    srv := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    log.Println("Serving gRPC-Gateway on http://localhost:8080")
    return srv.ListenAndServe()
}

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}