package usecases

import (
	"context"
	"event_service/internal/event/models"
	"event_service/internal/event/repositories"
)

type EventUseCase struct {
	repository *repositories.Repository
}

func NewEventUseCase(repo *repositories.Repository) *EventUseCase {
	return &EventUseCase{repository: repo}
}

func (c *EventUseCase) Create(ctx context.Context, createEvent *models.CreateEventInput) (*models.Event, error) {

	return nil, nil
}
