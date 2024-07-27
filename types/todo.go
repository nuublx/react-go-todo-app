package types

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateTodo struct {
	Description string `json:"description"`
}

type UpdateTodo struct {
	Description string `json:"description"`
}
