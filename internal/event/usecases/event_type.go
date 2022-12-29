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

func (c *EventTypeUseCase) List(ctx context.Context, filter *usecase_models.EventTypeFilter) (usecase_models.EventTypes, uint64, error) {
	// TODO: Validate

	types, count, err := c.repository.EventType.List(ctx, useCaseFilterEventTypeToRepo(filter))
	if err != nil {
		return nil, 0, err
	}

	return repoEventTypesDtoToUseCase(types), count, nil
}

func (c *EventTypeUseCase) Update(ctx context.Context, updateEventType *usecase_models.UpdateEventTypeInput) (*usecase_models.EventType, error) {
	// TODO: Validate

	updatedEventType, err := c.repository.EventType.Update(ctx, useCaseUpdateEventTypeDtoToRepo(updateEventType))
	if err != nil {
		return nil, err
	}

	return repoEventTypeDtoToUseCase(updatedEventType), nil
}

func (c *EventTypeUseCase) Delete(ctx context.Context, deleteEventType *usecase_models.DeleteEventTypeInput) error {

	return c.repository.EventType.Delete(ctx, useCaseDeleteEventTypeDtoToRepo(deleteEventType))
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

	repoFilter := &repository_models.EventTypeRepositoryFilter{
		Titles:     filter.Titles,
		Search:     filter.Search,
		PageSize:   filter.PageSize,
		PageNumber: filter.PageNumber,
	}

	if filter.OrderBy != nil {
		for _, useCaseOrderBy := range filter.OrderBy {
			switch useCaseOrderBy {
			case usecase_models.OrderByTypeCreatedAt:
				repoFilter.OrderBy = append(repoFilter.OrderBy, repository_models.OrderByTypeCreatedAt)
			case usecase_models.OrderByTypeTitle:
				repoFilter.OrderBy = append(repoFilter.OrderBy, repository_models.OrderByTypeTitle)
			case usecase_models.OrderByTypeID:
				repoFilter.OrderBy = append(repoFilter.OrderBy, repository_models.OrderByTypeID)
			}
		}
	}

	if filter.OrderDirection != nil {
		switch *filter.OrderDirection {
		case usecase_models.OrderDirectionASC:
			direction := repository_models.SortOrderASC
			repoFilter.SortOrder = &direction
		case usecase_models.OrderDirectionDESC:
			direction := repository_models.SortOrderDESC
			repoFilter.SortOrder = &direction
		}
	}

	return repoFilter
}

func repoEventTypesDtoToUseCase(types []*repository_models.EventTypeRepositoryDTO) usecase_models.EventTypes {
	useCaseTypes := make(usecase_models.EventTypes, 0, len(types))

	for _, t := range types {
		useCaseTypes = append(useCaseTypes, repoEventTypeDtoToUseCase(t))
	}

	return useCaseTypes
}

func useCaseUpdateEventTypeDtoToRepo(updateEventTypeUC *usecase_models.UpdateEventTypeInput) *repository_models.UpdateEventTypeRepositoryDTO {
	return &repository_models.UpdateEventTypeRepositoryDTO{
		Title:    updateEventTypeUC.Title,
		NewTitle: updateEventTypeUC.NewTitle,
	}
}

func useCaseDeleteEventTypeDtoToRepo(deleteEventTypeUC *usecase_models.DeleteEventTypeInput) *repository_models.DeleteEventTypeRepositoryDTO {
	return &repository_models.DeleteEventTypeRepositoryDTO{
		Title: deleteEventTypeUC.Title,
	}
}
