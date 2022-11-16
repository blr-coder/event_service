package pg_store

import (
	"context"
	"event_service/internal/event/models"
	"github.com/pkg/errors"

	// DB driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type EventTypePsqlStore struct {
	db *sqlx.DB
}

func NewEventTypePsqlStore(database *sqlx.DB) *EventTypePsqlStore {
	return &EventTypePsqlStore{db: database}
}

func (s *EventTypePsqlStore) Create(
	ctx context.Context,
	createEventType *models.CreateEventTypeInput,
) (*models.EventType, error) {

	query := `
		INSERT INTO event_types (title) 
		VALUES ($1)
		RETURNING id, title, created_at
	`

	eventType := &models.EventType{}

	return eventType, errors.WithStack(
		s.db.GetContext(ctx, eventType, query, createEventType.Title),
	)
}
