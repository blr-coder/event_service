package usecase_models

import (
	"event_service/internal/event/repositories/repository_models"
	"time"
)

type EventType struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateEventTypeInput struct {
	Title string `json:"title"`
}

type EventTypeFilter struct {
	Titles         []string                          `json:"titles"`
	Search         *string                           `json:"search"`
	OrderBy        repository_models.OrderByList     `json:"order_by"`
	OrderDirection *repository_models.OrderDirection `json:"order_direction"`
}

type EventTypes []*EventType

type UpdateEventTypeInput struct {
	Title    string `json:"title"`
	NewTitle string `json:"new_title"`
}

type DeleteEventTypeInput struct {
	Title string `json:"title"`
}
