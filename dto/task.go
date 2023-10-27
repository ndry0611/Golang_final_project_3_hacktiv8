package dto

import (
	"time"
)

type NewTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
}

type NewTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetTasksResponse struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Status      bool        `json:"status"`
	Description string      `json:"description"`
	UserID      uint        `json:"user_id"`
	CategoryID  uint        `json:"category_id"`
	CreatedAt   time.Time   `json:"created_at"`
	User        GetTaskUser `json:"User"`
}

type GetTaskUser struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type UpdateTaskRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateTaskStatusRequest struct {
	ID     uint `json:"id"`
	Status bool `json:"status" validate:"boolean"`
}

type UpdateTaskCategoryRequest struct {
	ID         uint `json:"id"`
	CategoryID uint `json:"category_id" validate:"required"`
}

type UpdateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
