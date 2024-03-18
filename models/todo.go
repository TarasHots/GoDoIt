package models

import "time"

type Todo struct {
	ID          string    `json:"id" validate:"required,number"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}
