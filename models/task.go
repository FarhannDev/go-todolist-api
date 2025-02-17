package models

type Task struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var Tasks []Task
