package routes

import (
	"goAuthTodo/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	//Auth Route
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//Todo Route
	app.Get("/api/todos", controllers.GetAllTodo)
	app.Get("/api/todo/:id", controllers.GetTodoByID)
	app.Post("/api/create", controllers.CreateTodo)
	app.Put("/api/update/:id", controllers.UpdateTodo)
	app.Delete("/api/delete/:id", controllers.DeleteTodo)
}
