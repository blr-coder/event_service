package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
)

type EventTypeUseCase struct {
	repository *repositories.Repository
}

func NewEventTypeUseCase(repo *repositories.Repository) *EventTypeUseCase {
	return &EventTypeUseCase{repository: repo}
}

func (c *EventTypeUseCase) Create(ctx context.Context, createEventType *usecase_models.CreateEventTypeInput) (*usecase_models.EventType, error) {
	// TODO: Validate

	eventType, err := c.repository.EventType.Create(ctx, useCaseCreateEventTypeDtoToRepo(createEventType))
	if err != nil {
		return nil, err
	}

	return RepoEventTypeDtoToUseCase(eventType), nil
}

func (c *EventTypeUseCase) Get(ctx context.Context, getEventType *usecase_models.GetEventTypeInput) (*usecase_models.EventType, error) {

	eventType, err := c.repository.EventType.Get(
		ctx,
		&repository_models.GetEventTypeRepositoryDTO{ID: getEventType.ID},
	)
	if err != nil {
		return nil, err
	}

	return RepoEventTypeDtoToUseCase(eventType), nil
}

func useCaseCreateEventTypeDtoToRepo(
	createEventTypeUC *usecase_models.CreateEventTypeInput,
) *repository_models.CreateEventTypeRepositoryDTO {
	return &repository_models.CreateEventTypeRepositoryDTO{Title: createEventTypeUC.Title}
}

func RepoEventTypeDtoToUseCase(repoEventType *repository_models.EventTypeRepositoryDTO) *usecase_models.EventType {
	return &usecase_models.EventType{
		ID:        repoEventType.ID,
		Title:     repoEventType.Title,
		CreatedAt: repoEventType.CreatedAt,
	}
}
