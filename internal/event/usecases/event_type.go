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

	return repoEventTypeDtoToUseCase(eventType), nil
}

func (c *EventTypeUseCase) List(ctx context.Context, filter *usecase_models.EventTypeFilter) (usecase_models.EventTypes, error) {
	// TODO: Validate

	types, err := c.repository.EventType.List(ctx, useCaseFilterEventTypeToRepo(filter))
	if err != nil {
		return nil, err
	}

	return repoEventTypesDtoToUseCase(types), nil
}

func useCaseCreateEventTypeDtoToRepo(
	createEventTypeUC *usecase_models.CreateEventTypeInput,
) *repository_models.CreateEventTypeRepositoryDTO {
	return &repository_models.CreateEventTypeRepositoryDTO{Title: createEventTypeUC.Title}
}

func repoEventTypeDtoToUseCase(repoEventType *repository_models.EventTypeRepositoryDTO) *usecase_models.EventType {
	return &usecase_models.EventType{
		Title:     repoEventType.Title,
		CreatedAt: repoEventType.CreatedAt,
		UpdatedAt: repoEventType.UpdatedAt,
	}
}

func useCaseFilterEventTypeToRepo(filter *usecase_models.EventTypeFilter) *repository_models.EventTypeRepositoryFilter {
	return &repository_models.EventTypeRepositoryFilter{
		Titles: filter.Titles,
		Search: filter.Search,
	}
}

func repoEventTypesDtoToUseCase(types []*repository_models.EventTypeRepositoryDTO) usecase_models.EventTypes {
	useCaseTypes := make(usecase_models.EventTypes, 0, len(types))

	for _, t := range types {
		useCaseTypes = append(useCaseTypes, repoEventTypeDtoToUseCase(t))
	}

	return useCaseTypes
}
