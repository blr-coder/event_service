package mongo_store

import (
	"context"
	"event_service/internal/event/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventTypeMongoStore struct {
	collection *mongo.Collection
}

func NewEventTypeMongoStore(mongoCollection *mongo.Collection) *EventTypeMongoStore {
	return &EventTypeMongoStore{collection: mongoCollection}
}

func (s *EventTypeMongoStore) Create(ctx context.Context, createEventType *models.CreateEventTypeInput) (*models.EventType, error) {

	return nil, nil
}
