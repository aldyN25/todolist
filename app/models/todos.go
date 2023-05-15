package models

import (
	"time"

	"github.com/aldyN25/todolist/pkg/utils/constants"
)

type (
	Todos struct {
		TodoId          uint `gorm:"column:todo_id;primary_key;auto_increment"`
		ActivityGroupId uint
		Title           string
		Priority        string
		IsActive        bool       `gorm:"column:is_active"`
		CreatedAt       time.Time  `gorm:"column:created_at"`
		UpdatedAt       *time.Time `gorm:"column:updated_at"`
		Activities      Activities `gorm:"foreignKey:ActivityGroupId;references:ActivityId"`
	}

	TodosRes struct {
		TodoId          uint   `json:"todo_id"`
		ActivityGroupId uint   `json:"activity_group_id"`
		Title           string `json:"title"`
		IsActive        bool   `json:"is_active"`
		Priority        string `json:"priority"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
	}

	TodosReq struct {
		Title           string `json:"title" validate:"required,min=3"`
		ActivityGroupId uint   `json:"activity_group_id" validate:"omitempty"`
		IsActive        bool   `json:"is_active"`
		Priority        string `json:"priority" validate:"omitempty"`
	}
)

func (Todos) TableName() string {
	return "todos"
}

func (n Todos) ToTodosRes() *TodosRes {
	result := &TodosRes{
		TodoId:          n.TodoId,
		ActivityGroupId: n.ActivityGroupId,
		Title:           n.Title,
		IsActive:        n.IsActive,
		Priority:        n.Priority,
		CreatedAt:       n.CreatedAt.Format(constants.LayoutYMD),
	}
	if n.UpdatedAt != nil {
		result.UpdatedAt = n.UpdatedAt.Format(constants.LayoutYMD)
	} else {
		result.UpdatedAt = ""
	}

	return result
}
