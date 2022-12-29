package usecases

import (
	"context"
	"event_service/internal/event/repositories"
	"event_service/internal/event/repositories/repository_models"
	"event_service/internal/event/usecases/usecase_models"
)

type EventUseCase struct {
	repository *repositories.Repository
}

func NewEventUseCase(repo *repositories.Repository) *EventUseCase {
	return &EventUseCase{repository: repo}
}

func (c *EventUseCase) Create(ctx context.Context, createEvent *usecase_models.CreateEventInput) (*usecase_models.Event, error) {
	// TODO: Validate

	event, err := c.repository.Event.Create(ctx, useCaseCreateEventDtoToRepo(createEvent))
	if err != nil {
		return nil, err
	}

	return repoEventDtoToUseCase(event), nil
}

func (c *EventUseCase) List(ctx context.Context, filter *usecase_models.EventFilter) (*usecase_models.Events, error) {
	// TODO: Validate

	events, count, err := c.repository.Event.List(ctx, useCaseEventFilterToRepo(filter))
	if err != nil {
		return nil, err
	}

	return repoEventsToUseCase(events, count), nil
}

func repoEventsToUseCase(repoEvents []*repository_models.EventRepositoryDTO, repoCount uint64) *usecase_models.Events {

	// TODO: Add loop

	return &usecase_models.Events{
		Events: nil,
		Count:  repoCount,
	}
}

func useCaseEventFilterToRepo(filter *usecase_models.EventFilter) *repository_models.EventRepositoryFilter {

	repoFilter := &repository_models.EventRepositoryFilter{
		TypeTitle:    filter.TypeTitle,
		CampaignID:   filter.CampaignID,
		InsertionID:  filter.InsertionID,
		UserID:       filter.UserID,
		CostCurrency: filter.CostCurrency,
		PageSize:     filter.PageSize,
		PageNumber:   filter.PageNumber,
	}

	if filter.SortBy != nil {
		for _, useCaseSortBy := range filter.SortBy {
			switch useCaseSortBy {
			case usecase_models.EventSortByCreatedAt:
				repoFilter.SortBy = append(repoFilter.SortBy, repository_models.EventSortByCreatedAt)
			case usecase_models.EventSortByTypeTitle:
				repoFilter.SortBy = append(repoFilter.SortBy, repository_models.EventSortByTypeTitle)
			case usecase_models.EventSortByID:
				repoFilter.SortBy = append(repoFilter.SortBy, repository_models.EventSortByID)
			case usecase_models.EventSortByCostCurrency:
				repoFilter.SortBy = append(repoFilter.SortBy, repository_models.EventSortByCostCurrency)
			}
		}
	}

	if filter.SortOrder != nil {
		switch *filter.SortOrder {
		case usecase_models.SortOrderASC:
			order := repository_models.SortOrderASC
			repoFilter.SortOrder = &order
		case usecase_models.SortOrderDESC:
			order := repository_models.SortOrderDESC
			repoFilter.SortOrder = &order
		}
	}

	return repoFilter
}

func useCaseCreateEventDtoToRepo(createEventUC *usecase_models.CreateEventInput) *repository_models.CreateEventRepositoryDTO {
	return &repository_models.CreateEventRepositoryDTO{
		TypeTitle:    createEventUC.TypeTitle,
		CampaignID:   createEventUC.CampaignID,
		InsertionID:  createEventUC.InsertionID,
		UserID:       createEventUC.UserID,
		CostAmount:   createEventUC.Cost.Amount,
		CostCurrency: createEventUC.Cost.Currency,
	}
}

func repoEventDtoToUseCase(repoEvent *repository_models.EventRepositoryDTO) *usecase_models.Event {
	return &usecase_models.Event{
		ID:           repoEvent.ID,
		TypeTitle:    repoEvent.TypeTitle,
		CampaignID:   repoEvent.CampaignID,
		InsertionID:  repoEvent.InsertionID,
		UserID:       repoEvent.UserID,
		CostAmount:   repoEvent.CostAmount,
		CostCurrency: repoEvent.CostCurrency,
		CreatedAt:    repoEvent.CreatedAt,
	}
}
