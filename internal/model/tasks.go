package model

import (
	"time"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null" validate:"required,min=1,max=255"`
	Description string     `json:"description" gorm:"type:text"`
	Completed   bool       `json:"completed" gorm:"default:false"`
	Priority    Priority   `json:"priority" gorm:"default:medium"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateTaskRequest struct {
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description,omitempty"`
	Priority    Priority   `json:"priority,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description *string    `json:"description,omitempty"`
	Completed   *bool      `json:"completed,omitempty"`
	Priority    *Priority  `json:"priority,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

type TaskResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Priority    Priority   `json:"priority"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (t *Task) ToResponse() *TaskResponse {
	return &TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		Priority:    t.Priority,
		Deadline:    t.Deadline,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
