package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/usecases/usecase_models"
)

type EventUseCase struct {
	repository *repositories.Repository
}

func NewEventUseCase(repo *repositories.Repository) *EventUseCase {
	return &EventUseCase{repository: repo}
}

func (c *EventUseCase) Create(ctx context.Context, createEvent *usecase_models.CreateEventInput) (*usecase_models.Event, error) {

	return nil, nil
}
