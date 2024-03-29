package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/usecases/usecase_models"
)

//go:generate mockgen -build_flags=-mod=mod -destination mock/event_type_mock.go event_service/internal/event/usecases IEventTypeUseCase

type IEventTypeUseCase interface {
	Create(ctx context.Context, createEventType *usecase_models.CreateEventTypeInput) (*usecase_models.EventType, error)
	List(ctx context.Context, filter *usecase_models.EventTypeFilter) (usecase_models.EventTypes, uint64, error)
	Update(ctx context.Context, updateEventType *usecase_models.UpdateEventTypeInput) (*usecase_models.EventType, error)
	Delete(ctx context.Context, deleteEventType *usecase_models.DeleteEventTypeInput) error
}

//go:generate mockgen -build_flags=-mod=mod -destination mock/event_mock.go event_service/internal/event/usecases IEventUseCase

type IEventUseCase interface {
	Create(ctx context.Context, createEvent *usecase_models.CreateEventInput) (*usecase_models.Event, error)
	List(ctx context.Context, filter *usecase_models.EventFilter) (*usecase_models.Events, error)
}

type IReportUseCase interface {
	List(ctx context.Context, filter *usecase_models.ReportFilter) (usecase_models.Reports, error)
}

type UseCase struct {
	EventType IEventTypeUseCase
	Event     IEventUseCase
	Report    IReportUseCase
}

func NewUseCase(repo *repositories.Repository) *UseCase {
	return &UseCase{
		EventType: NewEventTypeUseCase(repo),
		Event:     NewEventUseCase(repo),
		Report:    NewReportUseCase(repo),
	}
}
