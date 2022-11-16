package repositories

import (
	"context"
	"event_service/internal/event/models"
	"event_service/internal/event/repositories/mongo_store"
	"event_service/internal/event/repositories/pg_store"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEventTypeRepository interface {
	Create(ctx context.Context, createEventType *models.CreateEventTypeInput) (*models.EventType, error)
}

type IEventRepository interface {
	Create(ctx context.Context, createEvent *models.CreateEventInput) (*models.Event, error)
}

type Repository struct {
	EventType IEventTypeRepository
	Event     IEventRepository
}

func NewPsqlRepository(db *sqlx.DB) *Repository {
	return &Repository{
		EventType: pg_store.NewEventTypePsqlStore(db),
		Event:     pg_store.NewEventPsqlStore(db),
	}
}

// NewMongoRepository for example
func NewMongoRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		EventType: mongo_store.NewEventTypeMongoStore(collection),
		Event:     mongo_store.NewEventMongoStore(collection),
	}
}