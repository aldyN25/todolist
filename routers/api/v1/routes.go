package apiv1

import (
	activitiesController "github.com/aldyN25/todolist/app/controllers/activities"
	todosController "github.com/aldyN25/todolist/app/controllers/todo"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	router := app.Group("/api/v1")

	activities := router.Group("/activities-group")

	activities.Post("/", activitiesController.Createactivities)
	activities.Get("/", activitiesController.GetAllActivities)
	activities.Get("/:id", activitiesController.GetActivitiesById)
	activities.Put("/:id", activitiesController.UpdateActivity)
	activities.Delete("/:id", activitiesController.DeleteActivity)

	todos := router.Group("/todos-item")

	todos.Post("/", todosController.CreateTodos)
	todos.Get("/", todosController.GetAllTodos)
	todos.Get("/:id", todosController.GetTodosById)
	todos.Put("/:id", todosController.UpdateTodos)
	todos.Delete("/:id", todosController.DeleteTodos)
}
