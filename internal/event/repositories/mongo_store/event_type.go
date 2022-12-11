package mongo_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventTypeMongoStore struct {
	collection *mongo.Collection
}

func NewEventTypeMongoStore(mongoCollection *mongo.Collection) *EventTypeMongoStore {
	return &EventTypeMongoStore{collection: mongoCollection}
}

func (s *EventTypeMongoStore) Create(
	ctx context.Context,
	createEventType *repository_models.CreateEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	return nil, nil
}

func (s *EventTypeMongoStore) List(ctx context.Context, repositoryFilter *repository_models.EventTypeRepositoryFilter) ([]*repository_models.EventTypeRepositoryDTO, uint64, error) {

	return nil, 0, nil
}

func (s *EventTypeMongoStore) Update(
	ctx context.Context,
	updateEventType *repository_models.UpdateEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	return nil, nil
}

func (s *EventTypeMongoStore) Delete(
	ctx context.Context,
	deleteEventType *repository_models.DeleteEventTypeRepositoryDTO,
) error {

	return nil
}
