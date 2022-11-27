package repository_models

import "time"

type CreateEventTypeRepositoryDTO struct {
	Title string `db:"title"`
}

type EventTypeRepositoryDTO struct {
	ID        int64      `db:"id"`
	Title     string     `db:"title"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type GetEventTypeRepositoryDTO struct {
	ID int64 `json:"id"`
}
