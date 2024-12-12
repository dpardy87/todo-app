package models

type Todo struct {
	ID          string `json:"id"`          // Internal ID (within _source in Elasticsearch)
	Description string `json:"description"`
	TaskName    string `json:"task"`
	Completed   bool   `json:"completed"`
}
