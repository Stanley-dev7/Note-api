package main

import (
	"log"
	"net/http"

	"day33-notes-api/database"
	"day33-notes-api/handlers"
	"day33-notes-api/middleware"

	"github.com/gorilla/mux"
)

func main() {

	database.Connect()
	database.InitTables()

	r := mux.NewRouter()

	// PUBLIC ROUTES
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	}).Methods("GET")

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// PROTECTED ROUTES
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.ValidateToken)

	protected.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	protected.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	protected.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	protected.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}