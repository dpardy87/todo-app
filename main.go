package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"todo-app/api"
)

func main() {
	// Create router
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/api/todos", api.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", api.CreateTodo).Methods("POST")

	// For Prod: Serve static files from the generated Vue `/web` directory
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	// Setup CORS to allow connections from Vue dev server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Vite dev server origin
		AllowedMethods:   []string{"GET", "POST", "PUT"},   // Allowed HTTP methods
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Apply the CORS middleware to the router
	handler := c.Handler(router)

	// Start the server
	fmt.Println("Running Go server on port 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting Go server:", err)
	}
}

