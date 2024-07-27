package utils

import (
	"github.com/google/uuid"
	"github.com/nuublx/react-go-todo-app/types"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func FindTodoById(id string, todos *[]types.Todo) int {
	for index, todo := range *todos {
		if todo.ID == id {
			return index
		}
	}
	return -1
}
