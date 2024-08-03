package main

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	types "github.com/nuublx/react-go-todo-app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTodos(c *fiber.Ctx) error {
	var todos []types.Todo
	var err error
	todosCursor, err := todosCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	defer todosCursor.Close(context.Background())

	for todosCursor.TryNext(context.Background()) {
		var todo types.Todo
		if err = todosCursor.Decode(&todo); err != nil {
			return c.Status(500).JSON(fiber.Map{"err": true, "msg": "can't decode document into a todo"})
		}
		todos = append(todos, todo)
	}

	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{})
	}
	return c.Status(200).JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	var err error
	var result *mongo.InsertOneResult
	createTodoModel := new(types.CreateTodo)

	if err = c.BodyParser(createTodoModel); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if strings.TrimSpace(createTodoModel.Description) == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Description cannot be empty"})
	}

	newTodo := types.Todo{
		Completed:   false,
		Description: createTodoModel.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err = todosCollection.InsertOne(context.Background(), newTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}

	newTodo.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(newTodo)
}

func UpdateDescription(c *fiber.Ctx) error {
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	updateTodoModel := &types.UpdateTodo{}
	if err := c.BodyParser(updateTodoModel); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	update := bson.M{"$set": bson.M{"description": updateTodoModel.Description, "updatedAt": time.Now()}}

	var updatedTodo types.Todo
	err = todosCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectId}, update).Decode(&updatedTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	updatedTodo.Description = updateTodoModel.Description
	return c.Status(201).JSON(updatedTodo)
}

func MarkTodoCompleted(c *fiber.Ctx) error {
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	update := bson.M{"$set": bson.M{"completed": true, "updatedAt": time.Now()}}

	var todo types.Todo
	err = todosCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectId}, update).Decode(&todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	todo.Completed = true
	return c.Status(200).JSON(todo)
}

func MarkTodoPending(c *fiber.Ctx) error {
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	update := bson.M{"$set": bson.M{"completed": false, "updatedAt": time.Now()}}

	var todo types.Todo
	err = todosCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectId}, update).Decode(&todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	todo.Completed = false
	return c.Status(200).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	var todo types.Todo
	err = todosCollection.FindOneAndDelete(context.Background(), bson.M{"_id": objectId}).Decode(&todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	return c.Status(202).JSON(todo)
}
