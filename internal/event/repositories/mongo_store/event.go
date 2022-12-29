package mongo_store

import (
	"context"
	"event_service/internal/event/repositories/repository_models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventMongoStore struct {
	collection *mongo.Collection
}

func NewEventMongoStore(collection *mongo.Collection) *EventMongoStore {
	return &EventMongoStore{collection: collection}
}

func (s *EventMongoStore) Create(ctx context.Context, createEvent *repository_models.CreateEventRepositoryDTO) (*repository_models.EventRepositoryDTO, error) {

	return nil, nil
}

func (s *EventMongoStore) List(ctx context.Context, filter *repository_models.EventRepositoryFilter) ([]*repository_models.EventRepositoryDTO, uint64, error) {

	return nil, 0, nil
}
