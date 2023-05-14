package models

import (
	"time"

	"github.com/aldyN25/todolist/pkg/utils/constants"
)

type (
	Activities struct {
		ActivitiesId uint `gorm:"column:activities_id;primary_key;auto_increment"`
		Email        string
		Title        string
		CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt    *time.Time `gorm:"column:update_at"`
	}

	ActivitiesRes struct {
		ID        uint
		Email     string `json:"email"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"update_at"`
	}

	ActivitiesReq struct {
		Email string `json:"email" validate:"required,email"`
		Title string `json:"title" validate:"required"`
	}
)

func (activities Activities) ToActivitiesRes() *ActivitiesRes {
	result := &ActivitiesRes{
		ID:        activities.ActivitiesId,
		Email:     activities.Email,
		Title:     activities.Title,
		CreatedAt: activities.CreatedAt.Format(constants.LayoutYMD),
	}

	if activities.UpdatedAt != nil {
		result.UpdatedAt = activities.UpdatedAt.Format(constants.LayoutYMD)
	} else {
		result.UpdatedAt = ""
	}

	return result
}
