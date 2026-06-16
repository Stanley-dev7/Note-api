package main

import (
	"log"
	"net/http"
	"os"

	"day33-notes-api/database"
	"day33-notes-api/handlers"
	"day33-notes-api/middleware"

	"github.com/gorilla/mux"
)

func main() {

	// CONNECT DATABASE
	database.Connect()
	database.InitTables()

	// ROUTER
	r := mux.NewRouter()

	// HEALTH CHECK (for Render)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	}).Methods("GET")

	// PUBLIC ROUTES
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// PROTECTED ROUTES
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.ValidateToken)

	protected.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	protected.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	protected.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	protected.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

	// PORT CONFIG (RENDER FIX)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}