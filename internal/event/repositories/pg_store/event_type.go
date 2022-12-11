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
		RETURNING *
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
		SELECT title, created_at, updated_at FROM event_types WHERE deleted_at IS NULL
	`

	if repositoryFilter.Titles != nil {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("AND title IN ('%s')", strings.Join(repositoryFilter.Titles, "','")))
	}

	if repositoryFilter.Search != nil {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("AND title ILIKE '%%%s%%'", *repositoryFilter.Search))
	}

	if repositoryFilter.OrderBy != nil {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("ORDER BY (%s)", strings.Join(repositoryFilter.OrderBy.Join(), ",")))
	}

	if repositoryFilter.OrderDirection != nil && repositoryFilter.OrderBy != nil {
		query = fmt.Sprintf("%s %s", query, repositoryFilter.OrderDirection)
	}

	if repositoryFilter.OrderDirection != nil && repositoryFilter.OrderBy == nil {
		query = fmt.Sprintf("%s ORDER BY created_at %s", query, repositoryFilter.OrderDirection)
	}

	if repositoryFilter.PageSize != nil && repositoryFilter.PageNumber != nil {
		offset := *repositoryFilter.PageSize * (*repositoryFilter.PageNumber - 1)
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, *repositoryFilter.PageSize, offset)
	}

	return types, errors.WithStack(s.db.SelectContext(ctx, &types, query))
}

func (s *EventTypePsqlStore) Update(
	ctx context.Context,
	updateEventType *repository_models.UpdateEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	query := `
		UPDATE event_types SET title=$2, updated_at=(now() AT TIME ZONE 'utc') WHERE title=$1 RETURNING *
	`

	updatedEventType := &repository_models.EventTypeRepositoryDTO{}

	err := s.db.QueryRowxContext(ctx, query, updateEventType.Title, updateEventType.NewTitle).StructScan(updatedEventType)
	if err != nil {
		return nil, err
	}

	return updatedEventType, err
}

func (s *EventTypePsqlStore) Delete(
	ctx context.Context,
	deleteEventType *repository_models.DeleteEventTypeRepositoryDTO,
) error {

	query := `
		UPDATE event_types SET 
		                       updated_at=(now() AT TIME ZONE 'utc'), 
		                       deleted_at=(now() AT TIME ZONE 'utc') 
		                   WHERE title=$1 AND deleted_at IS NULL RETURNING *
	`

	res, err := s.db.ExecContext(ctx, query, deleteEventType.Title)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		// TODO: Add custom err
		return err
	}

	if affected == 0 {
		// TODO: Add custom err
		return errors.New("affected == 0")
	}

	return nil
}
