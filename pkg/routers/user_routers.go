package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuublx/react-go-todo-app/app/controllers"
)

func UsersRouters(app *fiber.App) {
	app.Post("/api/users/register", controllers.Register)
}
