package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todosCollection *mongo.Collection

func main() {
	fmt.Println("welcome to my to-do app")

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(fmt.Sprintf("error: %s", err.Error()))
	}
	Routers(app)

	dbURI := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	todosCollection = client.Database("todo-database").Collection("todos")

	PORT := os.Getenv("PORT")
	fmt.Printf("now listening on port: %s", PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
