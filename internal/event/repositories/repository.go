package repositories

import (
	"context"
	"event_service/internal/event/repositories/mongo_store"
	"event_service/internal/event/repositories/pg_store"
	"event_service/internal/event/repositories/repository_models"
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
	) ([]*repository_models.EventTypeRepositoryDTO, uint64, error)
	Update(
		ctx context.Context,
		updateEventType *repository_models.UpdateEventTypeRepositoryDTO,
	) (*repository_models.EventTypeRepositoryDTO, error)
	Delete(
		ctx context.Context,
		deleteEventType *repository_models.DeleteEventTypeRepositoryDTO,
	) error
}

//go:generate mockgen -build_flags=-mod=mod -destination mock/event_mock.go event_service/internal/event/repositories IEventRepository

type IEventRepository interface {
	Create(ctx context.Context, createEvent *repository_models.CreateEventRepositoryDTO) (*repository_models.EventRepositoryDTO, error)
	List(ctx context.Context, repositoryFilter *repository_models.EventRepositoryFilter) ([]*repository_models.EventRepositoryDTO, uint64, error)
}

type IReportRepository interface {
	List(ctx context.Context, repositoryFilter *repository_models.ReportRepositoryFilter) ([]*repository_models.Report, error)
}

type Repository struct {
	EventType IEventTypeRepository
	Event     IEventRepository
	Report    IReportRepository
}

func NewPsqlRepository(db *sqlx.DB) *Repository {
	return &Repository{
		EventType: pg_store.NewEventTypePsqlStore(db),
		Event:     pg_store.NewEventPsqlStore(db),
		Report:    pg_store.NewReportPsqlStore(db),
	}
}

// NewMongoRepository for example, we can switch between storages very quickly
func NewMongoRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		EventType: mongo_store.NewEventTypeMongoStore(collection),
		Event:     mongo_store.NewEventMongoStore(collection),
		//Report: .......
	}
}
