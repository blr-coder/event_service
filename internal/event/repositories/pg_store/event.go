package pg_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
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

	return nil, nil
}
