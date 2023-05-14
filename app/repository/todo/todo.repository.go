package todo

import (
	"time"

	"github.com/aldyN25/todolist/app/models"
	gormdb "github.com/aldyN25/todolist/pkg/gorm.db"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx, todos *[]models.Todos) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}
	err = db.Debug().Table("todos").Find(&todos).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTodosById(id string, todos *models.Todos) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("todos").Where("todos_id = ?", id).First(&todos).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateTodos(c *fiber.Ctx, todosRequest *models.TodosReq) (*models.Todos, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	todos := &models.Todos{
		ActivitiesId: todosRequest.ActivitiesId,
		Title:        todosRequest.Title,
		Priority:     todosRequest.Priority,
		IsActive:     todosRequest.IsActive,
		CreatedAt:    now,
	}

	err = db.Debug().Table("todos").Create(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func UpdateTodos(id string, todos *models.Todos, todosRequest *models.TodosReq) (*models.TodosRes, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	dataUpdates := models.Todos{
		Title:     todosRequest.Title,
		Priority:  todosRequest.Priority,
		IsActive:  todosRequest.IsActive,
		UpdatedAt: &now,
	}
	err = db.Debug().Table("todos").Where("todos_id = ?", id).Updates(&dataUpdates).Error
	if err != nil {
		return nil, err
	}

	err = db.Debug().Table("todos").First(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos.ToTodosRes(), nil
}

func DeleteTodos(id string, todos *models.Todos) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("todos").Delete(&todos).Error
	if err != nil {
		return err
	}
	return nil
}
