package repository_models

import "time"

type CreateEventTypeRepositoryDTO struct {
	Title string `db:"title"`
}

type EventTypeRepositoryDTO struct {
	Title     string     `db:"title"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type EventTypeRepositoryFilter struct {
	Titles []string `db:"titles"`
	Search *string  `db:"search"`
}

type UpdateEventTypeRepositoryDTO struct {
	Title    string `db:"title"`
	NewTitle string `db:"new_title"`
}
