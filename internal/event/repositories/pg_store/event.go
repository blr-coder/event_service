package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"github.com/pkg/errors"

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

func (s *EventPsqlStore) List(ctx context.Context, filter *repository_models.EventRepositoryFilter) ([]*repository_models.EventRepositoryDTO, uint64, error) {

	return nil, 0, nil
}
