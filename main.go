package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"todo-app/api"
)

func main() {
	// create router
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/api/todos", api.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", api.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", api.DeleteTodo).Methods("DELETE")

	// setup CORS to allow connections from Vue dev server
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"}, // Vite runs on 5173 by default
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // fetch, create, update, delete
        AllowedHeaders: []string{"Authorization", "Content-Type"},
        AllowCredentials: true,
    })

    // apply the CORS middleware to our router
    handler := c.Handler(router)

	// for Prod: serve static files from the generated vue /web directory
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	fmt.Println("Running server on port 8080...")
	// listen on 8080
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

