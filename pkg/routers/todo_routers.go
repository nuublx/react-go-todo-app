package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuublx/react-go-todo-app/app/controllers"
)

func TodoRouters(app *fiber.App) {

	// get all todos
	app.Get("/api/todos", controllers.GetAllTodos)

	// create a Todo
	app.Post("/api/todos", controllers.CreateTodo)

	// Update a Todo description
	app.Put("/api/todos/:id", controllers.UpdateDescription)

	// mark completed
	app.Put("/api/todos/mark-completed/:id", controllers.MarkTodoCompleted)

	// mark pending
	app.Put("/api/todos/mark-pending/:id", controllers.MarkTodoPending)

	// delete a Todo
	app.Delete("/api/todos/:id", controllers.DeleteTodo)

}
