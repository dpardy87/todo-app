package api

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
)

var todos = []Todo{}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = uuid.NewString()
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// func DeleteTodo(w http.ResponseWriter, r *http.Request){
// }

