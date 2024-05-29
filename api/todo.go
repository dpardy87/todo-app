package api

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	TaskName    string `json:"task"`
	Completed   bool   `json:"completed"`
}
