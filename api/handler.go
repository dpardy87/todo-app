package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"todo-app/elastic"
	"todo-app/models"
)

var todos = []models.Todo{} // slice literal

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Create an Elasticsearch client
	esHosts := []string{"http://localhost:9200"}
	esClient, err := elastic.NewClient(esHosts)
	if err != nil {
		http.Error(w, "Failed to create Elasticsearch client", http.StatusInternalServerError)
		return
	}

	// Perform a search query to get all todos from the "todos" index
	searchResult, err := esClient.GetAll(r.Context(), "todos")
	if err != nil {
		http.Error(w, "Failed to retrieve todos", http.StatusInternalServerError)
		return
	}

	if searchResult == nil || len(searchResult) == 0 {
		// return empty array
		searchResult = []models.Todo{}
	}

	// encode search result into response
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(searchResult); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// creating a todo

	// es config
	esHosts := []string{"http://localhost:9200"}
	esClient, err := elastic.NewClient(esHosts)
	if err != nil {
		http.Error(w, "Failed to create Elasticsearch client", http.StatusInternalServerError)
		return
	}

	// data to insert
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = uuid.New().String()

	id, err := esClient.Insert(r.Context(), "todos", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send todo to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	fmt.Println(`{"success": true, "id": "` + id + `"}`)
}