package usecase_models

import (
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
	Titles         []string        `json:"titles"`
	Search         *string         `json:"search"`
	OrderBy        []OrderBy       `json:"order_by"`
	OrderDirection *OrderDirection `json:"order_direction"`
	PageSize       *uint64         `json:"page_size"`
	PageNumber     *uint64         `json:"page_number"`
}

type OrderBy string

const (
	OrderByTypeCreatedAt OrderBy = "created_at"
	OrderByTypeTitle     OrderBy = "title"
	OrderByTypeID        OrderBy = "id"
)

type OrderDirection string

const (
	OrderDirectionASC  OrderDirection = "asc"
	OrderDirectionDESC OrderDirection = "desc"
)

type EventTypes []*EventType

type UpdateEventTypeInput struct {
	Title    string `json:"title"`
	NewTitle string `json:"new_title"`
}

type DeleteEventTypeInput struct {
	Title string `json:"title"`
}
