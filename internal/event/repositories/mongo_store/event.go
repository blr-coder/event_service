package mongo_store

import (
	"context"
	"event_service/internal/event/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventMongoStore struct {
	collection *mongo.Collection
}

func NewEventMongoStore(collection *mongo.Collection) *EventMongoStore {
	return &EventMongoStore{collection: collection}
}

func (s *EventMongoStore) Create(ctx context.Context, createEvent *models.CreateEventInput) (*models.Event, error) {

	return nil, nil
}
