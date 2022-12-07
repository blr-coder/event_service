package repositories

import (
	"context"
	"event_service/internal/event/repositories/mongo_store"
	"event_service/internal/event/repositories/pg_store"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEventTypeRepository interface {
	Create(
		ctx context.Context,
		createEventType *repository_models.CreateEventTypeRepositoryDTO,
	) (*repository_models.EventTypeRepositoryDTO, error)
	List(
		ctx context.Context,
		repositoryFilter *repository_models.EventTypeRepositoryFilter,
	) ([]*repository_models.EventTypeRepositoryDTO, error)
	Update(
		ctx context.Context,
		updateEventType *repository_models.UpdateEventTypeRepositoryDTO,
	) (*repository_models.EventTypeRepositoryDTO, error)
	Delete(
		ctx context.Context,
		deleteEventType *repository_models.DeleteEventTypeRepositoryDTO,
	) error
}

type IEventRepository interface {
	Create(ctx context.Context, createEvent *usecase_models.CreateEventInput) (*usecase_models.Event, error)
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

// NewMongoRepository for example, we can switch between storages very quickly
func NewMongoRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		EventType: mongo_store.NewEventTypeMongoStore(collection),
		Event:     mongo_store.NewEventMongoStore(collection),
	}
}
