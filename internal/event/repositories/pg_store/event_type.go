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
) (types []*repository_models.EventTypeRepositoryDTO, count uint64, err error) {

	listQuery := `
		SELECT title, created_at, updated_at FROM event_types WHERE deleted_at IS NULL
	`
	countQuery := `
		SELECT count(*) FROM event_types WHERE deleted_at IS NULL
	`

	listTypes, args := s.decodeRepositoryFilter(listQuery, repositoryFilter, true)
	countTypes, args := s.decodeRepositoryFilter(countQuery, repositoryFilter, false)

	err = errors.WithStack(s.db.SelectContext(ctx, &types, listTypes, args...))
	if err != nil {
		return nil, 0, err
	}
	err = errors.WithStack(s.db.GetContext(ctx, &count, countTypes, args...))
	if err != nil {
		return nil, 0, err
	}

	return types, count, nil
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

func (s *EventTypePsqlStore) decodeRepositoryFilter(
	query string,
	filter *repository_models.EventTypeRepositoryFilter,
	paginate bool,
) (string, []any) {

	query = fmt.Sprintf("%s", query)
	var args []any

	if filter.Titles != nil {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("AND title IN ('%s')", strings.Join(filter.Titles, "','")))
	}

	if filter.Search != nil {
		search := fmt.Sprintf("%%%s%%", *filter.Search)
		query = fmt.Sprintf("%s AND title ILIKE ?", query)
		args = append(args, search)
	}

	if filter.OrderBy != nil && paginate {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("ORDER BY (%s)", strings.Join(filter.OrderBy.Join(), ",")))
	}
	if filter.SortOrder != nil && filter.OrderBy != nil && paginate {
		query = fmt.Sprintf("%s %s", query, filter.SortOrder)
	}

	if paginate && filter.PageSize != nil && filter.PageNumber != nil {
		if *filter.PageNumber == 0 {
			*filter.PageNumber = 1
		}
		offset := *filter.PageSize * (*filter.PageNumber - 1)
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, *filter.PageSize, offset)
	}

	query = s.db.Rebind(query)
	return query, args
}
