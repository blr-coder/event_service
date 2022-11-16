package usecases

import (
	"context"
	"event_service/internal/event/models"
	"event_service/internal/event/repositories"
)

type EventTypeUseCase struct {
	repository *repositories.Repository
}

func NewEventTypeUseCase(repo *repositories.Repository) *EventTypeUseCase {
	return &EventTypeUseCase{repository: repo}
}

func (c *EventTypeUseCase) Create(ctx context.Context, createEventType *models.CreateEventTypeInput) (*models.EventType, error) {

	return nil, nil
}
