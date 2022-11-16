package models

import "time"

type EventType struct {
	ID        int64      `db:"id"`
	Title     string     `db:"title"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type CreateEventTypeInput struct {
	Title string `db:"title"`
}
