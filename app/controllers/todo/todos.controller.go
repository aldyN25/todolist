package todo

import (
	"fmt"

	"github.com/aldyN25/todolist/app/models"
	todosRepository "github.com/aldyN25/todolist/app/repository/todo"
	"github.com/aldyN25/todolist/pkg/utils/constants"
	"github.com/aldyN25/todolist/pkg/utils/converter"
	"github.com/aldyN25/todolist/pkg/utils/validator"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	var todos []models.Todos
	err := todosRepository.GetAllTodos(c, &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "Internal Service Error",
			"error":   err.Error(),
		})
	}

	todosRes := converter.MapTodosToTodosRes(todos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "Ok",
		"data":    todosRes,
	})
}

func GetTodosById(c *fiber.Ctx) error {
	id := c.Params("id")
	todos := models.Todos{}

	err := todosRepository.GetTodosById(id, &todos)
	if err != nil {
		msgErr := fmt.Sprintf("Todo with ID %v Not Found", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Not Found",
			"message": msgErr,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    todos.ToTodosRes(),
	})
}

func CreateTodos(c *fiber.Ctx) error {
	todosRequest := new(models.TodosReq)

	err := c.BodyParser(&todosRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": err.Error(),
		})
	}

	errors := validator.ValidateRequest(todosRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	note, _ := todosRepository.CreateTodos(c, todosRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    note.ToTodosRes(),
	})
}

func UpdateTodos(c *fiber.Ctx) error {
	todosRequest := new(models.TodosReq)

	c.BodyParser(&todosRequest)

	errors := validator.ValidateRequest(todosRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"errors":  errors,
		})
	}

	id := c.Params("id")
	todos := models.Todos{}
	err := todosRepository.GetTodosById(id, &todos)

	if err != nil {
		msgErr := fmt.Sprintf("Todo with ID %v Not Found", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Not Found",
			"message": msgErr,
			"error":   err.Error(),
		})
	}

	todosRepository.UpdateTodos(id, &todos, todosRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    todos.ToTodosRes(),
	})
}

func DeleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")

	todos := models.Todos{}
	err := todosRepository.GetTodosById(id, &todos)

	if err != nil {
		msgErr := fmt.Sprintf("Todo with ID %v Not Found", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Not Found",
			"message": msgErr,
			"error":   err.Error(),
		})
	}

	todosRepository.DeleteTodos(id, &todos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "Success",
		"data":    nil,
	})
}
