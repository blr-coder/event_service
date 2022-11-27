package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
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
	createEventType *repository_models.CreateEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	query := `
		INSERT INTO event_types (title) 
		VALUES ($1)
		RETURNING id, title, created_at
	`

	eventType := &repository_models.EventTypeRepositoryDTO{}

	return eventType, errors.WithStack(
		s.db.GetContext(ctx, eventType, query, createEventType.Title),
	)
}

func (s *EventTypePsqlStore) Get(
	ctx context.Context,
	getEventType *repository_models.GetEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	query := `
		SELECT id, title, created_at, deleted_at 
		FROM event_types 
		WHERE id=$1 AND deleted_at IS NULL
	`

	eventType := &repository_models.EventTypeRepositoryDTO{}

	return eventType, errors.WithStack(
		s.db.GetContext(ctx, eventType, query, getEventType.ID),
	)
}
