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

func (s *EventTypeMongoStore) Get(
	ctx context.Context,
	getEventType *repository_models.GetEventTypeRepositoryDTO,
) (*repository_models.EventTypeRepositoryDTO, error) {

	return nil, nil
}
