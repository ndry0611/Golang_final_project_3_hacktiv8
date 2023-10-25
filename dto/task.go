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
