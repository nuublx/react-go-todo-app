package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nuublx/react-go-todo-app/pkg/routers"
	mongodb "github.com/nuublx/react-go-todo-app/platform/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var todosCollection *mongo.Collection

func main() {
	fmt.Println("welcome to my to-do app")

	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(fmt.Sprintf("error: %s", err.Error()))
	}
	routers.TodoRouters(app)
	routers.UsersRouters(app)
	mongodb.OpenMongoConnection()

	defer mongodb.Client.Disconnect(context.Background())

	PORT := os.Getenv("PORT")
	fmt.Printf("now listening on port: %s", PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
