package app

import (
	"fmt"
	"login-go/infrastructure"
	"login-go/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	// 1. Define handlers with their service and repository
	userHandler := NewUserHandler(service.NewUserService(infrastructure.NewUserRepositoryStub()))
	// 2. Create routes
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
	// 3. Create server
	fmt.Println("listen on 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err)
	}
}
