package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"todo-app/elastic"
)

var todos = []Todo{} // slice literal

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// creating a todo

	// es config
	esHosts := []string{"http://localhost:9200"}
	esClient := elastic.NewClient(esHosts)
	if esClient == nil {
		http.Error(w, "Failed to create Elasticsearch client", http.StatusInternalServerError)
		return
	}

	// data to insert
	var todo Todo
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
	json.NewEncoder(w).Encode(todo)
	fmt.Println(`{"success": true, "id": "` + id + `"}`)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // fetch route vars
	id := vars["id"]    // get value of 'id' from URL
	for index, todo := range todos {
		if todo.ID == id {
			/*
				- todos[:index] creates new slice including elements up to but not including, the element at index
				- todos[index+1:] creates a slice that starts just after the element at index (goes to end of slice)
				- ... combines the two slices and returns the value, assigned to todos
			*/
			todos = append(todos[:index], todos[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// if not found
	w.WriteHeader(http.StatusNotFound)
	// map[KeyType]ValueType
	json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// updating a todo
	vars := mux.Vars(r) // fetch route variables
	id := vars["id"]    // get value of 'id' from URL
	var updatedTodo Todo

	// create new decoder, read from reqeust body, and pass that data to updatedTodo (by ref)
	json.NewDecoder(r.Body).Decode(&updatedTodo)

	for index, todo := range todos {
		if todo.ID == id {
			todos[index].TaskName = updatedTodo.TaskName
			todos[index].Description = updatedTodo.Description
			todos[index].Completed = updatedTodo.Completed

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// serialize the todos[index] item, returns a new Encoder that writes to w
			json.NewEncoder(w).Encode(todos[index])
		}
	}
}
