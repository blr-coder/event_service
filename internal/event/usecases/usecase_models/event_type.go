package usecase_models

import "time"

type EventType struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateEventTypeInput struct {
	Title string `json:"title"`
}

type GetEventTypeInput struct {
	ID int64 `json:"id"`
}
