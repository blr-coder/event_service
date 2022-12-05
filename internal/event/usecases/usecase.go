package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/usecases/usecase_models"
)

type IEventTypeUseCase interface {
	Create(ctx context.Context, createEventType *usecase_models.CreateEventTypeInput) (*usecase_models.EventType, error)
	List(ctx context.Context, filter *usecase_models.EventTypeFilter) (usecase_models.EventTypes, error)
	Update(ctx context.Context, updateEventType *usecase_models.UpdateEventTypeInput) (*usecase_models.EventType, error)
}

type IEventUseCase interface {
	Create(ctx context.Context, createEvent *usecase_models.CreateEventInput) (*usecase_models.Event, error)
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
