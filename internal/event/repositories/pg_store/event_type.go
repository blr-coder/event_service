package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"fmt"
	"github.com/pkg/errors"
	"strings"

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
		RETURNING title, created_at, updated_at
	`

	eventType := &repository_models.EventTypeRepositoryDTO{}

	return eventType, errors.WithStack(
		s.db.GetContext(ctx, eventType, query, createEventType.Title),
	)
}

func (s *EventTypePsqlStore) List(
	ctx context.Context,
	repositoryFilter *repository_models.EventTypeRepositoryFilter,
) (types []*repository_models.EventTypeRepositoryDTO, err error) {

	query := `
		SELECT title, created_at, updated_at FROM event_types WHERE true
	`

	if repositoryFilter.Titles != nil {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("AND title IN ('%s')", strings.Join(repositoryFilter.Titles, "','")))
	}

	if repositoryFilter.Search != nil {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("AND title ILIKE '%%%s%%'", *repositoryFilter.Search))
	}

	// TODO: Add orderBy

	return types, errors.WithStack(s.db.SelectContext(ctx, &types, query))
}

func (s *EventTypePsqlStore) Update(
	ctx context.Context,
	updateEventType *repository_models.UpdateEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	query := `
		UPDATE event_types SET title=$2, updated_at=(now() AT TIME ZONE 'utc') WHERE title=$1 RETURNING *
	`

	//_, err := s.db.ExecContext(ctx, query, updateEventType.Title, updateEventType.NewTitle)
	_, err := s.db.QueryxContext(ctx, query, updateEventType.Title, updateEventType.NewTitle)
	if err != nil {
		return nil, err
	}

	// TODO: Add scan to struct
	return nil, err
}
