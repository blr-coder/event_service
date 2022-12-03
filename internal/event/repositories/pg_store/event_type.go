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
) ([]*repository_models.EventTypeRepositoryDTO, error) {

	query := `
		SELECT title, created_at, updated_at FROM event_types WHERE true
	`

	// TODO: Need better solution
	if repositoryFilter.Titles != nil {
		query = fmt.Sprintf("%s %s", query, "AND title IN ('"+strings.Join(repositoryFilter.Titles, "','")+"')")
	}

	var types []*repository_models.EventTypeRepositoryDTO

	return types, errors.WithStack(s.db.SelectContext(ctx, &types, query))
}
