package main

import (
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {

	// get all todos
	app.Get("/api/todos", GetAllTodos)

	// create a Todo
	app.Post("/api/todos", CreateTodo)

	// Update a Todo description
	app.Put("/api/todos/:id", UpdateDescription)

	// mark completed
	app.Put("/api/todos/mark-completed/:id", MarkTodoCompleted)

	// mark pending
	app.Put("/api/todos/mark-pending/:id", MarkTodoPending)

	// delete a Todo
	app.Delete("/api/todos/:id", DeleteTodo)

}
