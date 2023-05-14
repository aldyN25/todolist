package models

import (
	"time"

	"github.com/aldyN25/todolist/pkg/utils/constants"
)

type (
	Todos struct {
		TodosId      uint `gorm:"column:todos_id;primary_key;auto_increment"`
		ActivitiesId uint
		Title        string
		Priority     string
		IsActive     bool       `gorm:"column:is_active"`
		CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt    *time.Time `gorm:"column:updated_at"`
	}

	TodosRes struct {
		ID           uint   `json:"id"`
		ActivitiesId uint   `json:"activities_group_id"`
		Title        string `json:"title"`
		IsActive     bool   `json:"is_active"`
		Priority     string `json:"priority"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	TodosReq struct {
		Title        string `json:"title" validate:"required,min=3"`
		ActivitiesId uint   `json:"activities_id" validate:"omitempty"`
		IsActive     bool   `json:"is_active"`
		Priority     string `json:"priority" validate:"omitempty"`
	}
)

func (n Todos) ToTodosRes() *TodosRes {
	result := &TodosRes{
		ID:           n.TodosId,
		ActivitiesId: n.ActivitiesId,
		Title:        n.Title,
		IsActive:     n.IsActive,
		Priority:     n.Priority,
		CreatedAt:    n.CreatedAt.Format(constants.LayoutYMD),
	}
	if n.UpdatedAt != nil {
		result.UpdatedAt = n.UpdatedAt.Format(constants.LayoutYMD)
	} else {
		result.UpdatedAt = ""
	}

	return result
}
