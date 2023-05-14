package activities

import (
	"log"
	"time"

	"github.com/aldyN25/todolist/app/models"
	gormdb "github.com/aldyN25/todolist/pkg/gorm.db"
	"github.com/gofiber/fiber/v2"
)

func GetAllActivity(c *fiber.Ctx, activities *[]models.Activities) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}
	err = db.Debug().Table("activities").Find(&activities).Error
	if err != nil {
		return err
	}
	return nil
}

func GetActivityById(id string, activities *models.Activities) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("activities").First(&activities, id).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateActivity(c *fiber.Ctx, activitiesRequest *models.ActivitiesReq) (*models.Activities, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}

	activities := &models.Activities{
		Email:     activitiesRequest.Email,
		Title:     activitiesRequest.Title,
		CreatedAt: time.Now(),
	}

	err = db.Debug().Table("activities").Create(&activities).Error
	if err != nil {
		return nil, err
	}

	log.Println("log activities : ", activities)
	return activities, nil
}

func UpdateActivity(id string, activities *models.Activities, activitiesRequest *models.ActivitiesReq) (*models.ActivitiesRes, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	dataUpdates := models.Activities{
		Email:     activitiesRequest.Email,
		Title:     activitiesRequest.Title,
		UpdatedAt: &now,
	}
	err = db.Debug().Table("activities").Where("activities_id = ?", id).Updates(&dataUpdates).Error
	if err != nil {
		return nil, err
	}

	err = db.Debug().Table("activities").Where("activities_id = ? ", id).First(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities.ToActivitiesRes(), err
}

func DeleteActivity(id string, activities *models.Activities) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("activities").Where("activities_id = ?", id).Delete(&activities).Error
	if err != nil {
		return err
	}
	return nil
}
