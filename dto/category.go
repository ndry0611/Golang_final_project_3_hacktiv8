package dto

import "time"

type NewCategoryRequest struct {
	Type string `json:"type" validate:"required"`
}

type NewCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
