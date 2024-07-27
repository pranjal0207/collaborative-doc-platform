package main

import (
	"log"
	"net/http"

	userServiceHandle "pranjal0207/collaborative-doc-platform/api-gateway/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/registerUser", userServiceHandle.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", userServiceHandle.LoginUserHandler).Methods("POST")
	r.HandleFunc("/user", userServiceHandle.GetUserProfileHandler).Methods("GET")

	log.Printf("API Gateway listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
