package main

import (
	"log"
	"net/http"
	"task-manager/config"
	"task-manager/handlers"
	"task-manager/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	taskRouter := r.PathPrefix("/tasks").Subrouter()
	taskRouter.Use(middleware.JWTMiddleware)
	taskRouter.HandleFunc("", handlers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("", handlers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods("DELETE")

	// Middleware for CORS and logging
	handler := middleware.CORS(middleware.Logger(r))

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
