package activities

import (
	"github.com/aldyN25/todolist/app/models"
	activitiesRepository "github.com/aldyN25/todolist/app/repository/activities"
	"github.com/aldyN25/todolist/pkg/utils/constants"
	"github.com/aldyN25/todolist/pkg/utils/converter"
	"github.com/aldyN25/todolist/pkg/utils/validator"
	"github.com/gofiber/fiber/v2"
)

func GetAllActivities(c *fiber.Ctx) error {
	var activities []models.Activities
	err := activitiesRepository.GetAllActivity(c, &activities)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "Internal Service Error",
			"error":   err.Error(),
		})
	}

	activitiesRes := converter.MapActivitiesToActivitiesRes(activities)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    activitiesRes,
	})
}

func GetActivitiesById(c *fiber.Ctx) error {
	id := c.Params("id")
	activities := models.Activities{}

	err := activitiesRepository.GetActivityById(id, &activities)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "activities not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "activities found",
		"data":    activities.ToActivitiesRes(),
	})
}

func Createactivities(c *fiber.Ctx) error {
	activitiesRequest := new(models.ActivitiesReq)

	err := c.BodyParser(&activitiesRequest)
	if err != nil {
		return err
	}

	errors := validator.ValidateRequest(activitiesRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	activities, _ := activitiesRepository.CreateActivity(c, activitiesRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    activities.ToActivitiesRes(),
	})
}

func UpdateActivity(c *fiber.Ctx) error {
	activitiesRequest := new(models.ActivitiesReq)

	c.BodyParser(&activitiesRequest)

	errors := validator.ValidateRequest(activitiesRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"errors":  errors,
		})
	}

	id := c.Params("id")
	activities := models.Activities{}
	err := activitiesRepository.GetActivityById(id, &activities)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "activities not found",
		})
	}

	activitiesRepository.UpdateActivity(id, &activities, activitiesRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "activities updated",
		"data":    activities.ToActivitiesRes(),
	})
}

func DeleteActivity(c *fiber.Ctx) error {
	id := c.Params("id")

	activities := models.Activities{}
	err := activitiesRepository.GetActivityById(id, &activities)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "activities not found",
		})
	}

	activitiesRepository.DeleteActivity(id, &activities)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "Success",
		"data":    nil,
	})
}
