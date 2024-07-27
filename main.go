package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	types "github.com/nuublx/react-go-todo-app/types"
	"github.com/nuublx/react-go-todo-app/utils"
)

var todos []types.Todo

func main() {
	fmt.Println("welcome to my to-do app")

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(fmt.Sprintf("error: %s", err.Error()))
	}

	// get all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		if len(todos) <= 0 {
			err := errors.New("there is no todo items yet")

			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(todos)
	})

	// create a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		createTodoModel := &types.CreateTodo{}
		if err := c.BodyParser(createTodoModel); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		if strings.TrimSpace(createTodoModel.Description) == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Description cannot be empty"})
		}

		newTodo := types.Todo{
			ID:          utils.GenerateUUID(),
			Completed:   false,
			Description: createTodoModel.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		todos = append(todos, newTodo)

		return c.Status(201).JSON(newTodo)
	})

	// Update a Todo description
	app.Put("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		updateTodoModel := &types.UpdateTodo{}
		if err := c.BodyParser(updateTodoModel); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		todoIndex := utils.FindTodoById(id, &todos)
		if todoIndex == -1 {
			return c.Status(400).JSON(fiber.Map{"error": "Todo does not exist"})
		}

		todoToUpdate := &todos[todoIndex]
		todoToUpdate.Description = updateTodoModel.Description
		todoToUpdate.UpdatedAt = time.Now()

		return c.Status(201).JSON(todoToUpdate)
	})

	// toggle completed/ not-completed
	app.Put("/api/todos/toggle-completion/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		todoIndex := utils.FindTodoById(id, &todos)
		if todoIndex == -1 {
			return c.Status(400).JSON(fiber.Map{"error": "Todo does not exist"})
		}

		todo := &todos[todoIndex]
		todo.Completed = !(*todo).Completed
		todo.UpdatedAt = time.Now()

		return c.Status(200).JSON(todo)
	})

	// delete a Todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		todoIndex := utils.FindTodoById(id, &todos)
		if todoIndex == -1 {
			return c.Status(400).JSON(fiber.Map{"error": "Todo does not exist"})
		}

		todos = append(todos[:todoIndex], todos[todoIndex+1:]...)

		return c.SendStatus(202)
	})

	PORT := os.Getenv("PORT")
	fmt.Printf("now listening on port: %s", PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
