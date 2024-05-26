package api

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var todos = []Todo{} // slice literal

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// creating a todo
	var todo Todo
	
	// create a new decoder, read from request body,
	// decodes the JSON data into todo (passed by reference)
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = uuid.NewString()

	// adds todo to todos slice
	todos = append(todos, todo)

	// creates a new encoder: writes to the response writer (w)
	// encodes the todo struct as JSON, writes it to HTTP response
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)	// fetch route vars
	id := vars["id"] 	// get value of 'id' from URL
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

func UpdateTodo(w http.ResponseWriter, r *http.Request){
	// updating a todo
	vars := mux.Vars(r) // fetch route variables
	id := vars["id"]	// get value of 'id' from URL
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
