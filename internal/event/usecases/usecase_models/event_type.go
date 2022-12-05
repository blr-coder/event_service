package usecase_models

import "time"

type EventType struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateEventTypeInput struct {
	Title string `json:"title"`
}

type EventTypeFilter struct {
	Titles []string `json:"titles"`
	Search *string  `json:"search"`
}

type EventTypes []*EventType

type UpdateEventTypeInput struct {
	Title    string `json:"title"`
	NewTitle string `json:"new_title"`
}
