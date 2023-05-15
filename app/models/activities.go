package models

import (
	"time"

	"github.com/aldyN25/todolist/pkg/utils/constants"
)

type (
	Activities struct {
		ActivityId uint `gorm:"column:activity_id;primary_key;auto_increment"`
		Email      string
		Title      string
		CreatedAt  time.Time  `gorm:"column:created_at"`
		UpdatedAt  *time.Time `gorm:"column:updated_at"`
	}

	ActivitiesRes struct {
		ActivityId uint   `json:"activity_id"`
		Email      string `json:"email"`
		Title      string `json:"title"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"update_at"`
	}

	ActivitiesReq struct {
		Email string `json:"email" validate:"required,email"`
		Title string `json:"title" validate:"required"`
	}
)

func (Activities) TableName() string {
	return "activities"
}

func (activities Activities) ToActivitiesRes() *ActivitiesRes {
	result := &ActivitiesRes{
		ActivityId: activities.ActivityId,
		Email:      activities.Email,
		Title:      activities.Title,
		CreatedAt:  activities.CreatedAt.Format(constants.LayoutYMD),
	}

	if activities.UpdatedAt != nil {
		result.UpdatedAt = activities.UpdatedAt.Format(constants.LayoutYMD)
	} else {
		result.UpdatedAt = ""
	}

	return result
}
