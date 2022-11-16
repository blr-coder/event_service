package usecases

import (
	"context"
	"event_service/internal/event/models"
	"event_service/internal/event/repositories"
)

type IEventTypeUseCase interface {
	Create(ctx context.Context, createEventType *models.CreateEventTypeInput) (*models.EventType, error)
}

type IEventUseCase interface {
	Create(ctx context.Context, createEvent *models.CreateEventInput) (*models.Event, error)
}

type UseCase struct {
	EventType IEventTypeUseCase
	Event     IEventUseCase
}

func NewUseCase(repo *repositories.Repository) *UseCase {
	return &UseCase{
		EventType: NewEventTypeUseCase(repo),
		Event:     NewEventUseCase(repo),
	}
}
