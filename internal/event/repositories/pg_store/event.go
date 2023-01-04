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

type EventPsqlStore struct {
	db *sqlx.DB
}

func NewEventPsqlStore(database *sqlx.DB) *EventPsqlStore {
	return &EventPsqlStore{db: database}
}

func (s *EventPsqlStore) Create(
	ctx context.Context, createEvent *repository_models.CreateEventRepositoryDTO,
) (*repository_models.EventRepositoryDTO, error) {

	query := `
		INSERT INTO events 
		    (type_title, campaign_id, insertion_id, user_id, cost_amount, cost_currency)
		VALUES 
		    ($1, $2, $3, $4, $5, $6)
		RETURNING *
`

	event := &repository_models.EventRepositoryDTO{}

	return event, errors.WithStack(
		s.db.GetContext(
			ctx,
			event,
			query,
			createEvent.TypeTitle,
			createEvent.CampaignID,
			createEvent.InsertionID,
			createEvent.UserID,
			createEvent.CostAmount,
			createEvent.CostCurrency,
		),
	)
}

func (s *EventPsqlStore) List(ctx context.Context, filter *repository_models.EventRepositoryFilter) (events []*repository_models.EventRepositoryDTO, count uint64, err error) {

	listQuery := `
		SELECT id, type_title, campaign_id, insertion_id, user_id, cost_amount, cost_currency, created_at FROM events
	`
	countQuery := `
		SELECT count(*) FROM events
	`

	listEvents, args := s.decodeRepositoryFilter(listQuery, filter, true)
	countEvents, args := s.decodeRepositoryFilter(countQuery, filter, false)

	err = errors.WithStack(s.db.SelectContext(ctx, &events, listEvents, args...))
	if err != nil {
		return nil, 0, err
	}
	err = errors.WithStack(s.db.GetContext(ctx, &count, countEvents, args...))
	if err != nil {
		return nil, 0, err
	}

	return events, count, nil
}

func (s *EventPsqlStore) decodeRepositoryFilter(
	query string,
	filter *repository_models.EventRepositoryFilter,
	paginate bool,
) (string, []any) {

	query = fmt.Sprintf("%s WHERE TRUE", query)
	var args []any

	if filter.TypeTitle != nil {
		query = fmt.Sprintf("%s AND type_title =(?)", query)
		args = append(args, filter.TypeTitle)
	}
	if filter.CampaignID != nil {
		query = fmt.Sprintf("%s AND campaign_id =(?)", query)
		args = append(args, filter.CampaignID)
	}
	if filter.InsertionID != nil {
		query = fmt.Sprintf("%s AND insertion_id =(?)", query)
		args = append(args, filter.InsertionID)
	}
	if filter.UserID != nil {
		query = fmt.Sprintf("%s AND user_id =(?)", query)
		args = append(args, filter.UserID)
	}
	if filter.CostCurrency != nil {
		query = fmt.Sprintf("%s AND cost_currency =(?)", query)
		args = append(args, filter.CostCurrency)
	}

	if filter.SortBy != nil && paginate {
		query = fmt.Sprintf("%s %s", query, fmt.Sprintf("ORDER BY (%s)", strings.Join(filter.SortBy.Join(), ",")))
	}
	if filter.SortOrder != nil && filter.SortBy != nil && paginate {
		query = fmt.Sprintf("%s %s", query, filter.SortOrder)
	}

	// TODO: Add other fields check

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
